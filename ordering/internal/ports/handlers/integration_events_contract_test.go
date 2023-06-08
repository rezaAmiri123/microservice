///go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"github.com/pact-foundation/pact-go/v2/matchers"
	v4 "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type String = matchers.String
type Map = matchers.Map

var Decimal = matchers.Decimal
var Like = matchers.Like
var EachLike = matchers.EachLike
var ArrayContaining = matchers.ArrayContaining

func TestOrdersConsumer(t *testing.T) {
	type mocks struct {
		app *app.MockApp
		//orders *domain.MockOrderRepository
	}

	type rawEvent struct {
		Name    string
		Payload map[string]any
	}

	reg := registry.New()
	serde := serdes.NewJsonSerde(reg)
	err := orderingpb.RegistrationsWithSerde(serde)
	if err != nil {
		t.Fatal(err)
	}
	err = basketspb.RegistrationWithSerde(serde)
	if err != nil {
		t.Fatal(err)
	}

	err = depotpb.RegistrationsWithSerde(serde)
	if err != nil {
		t.Fatal(err)
	}

	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "ordering-pub",
		Consumer: "ordering-sub",
		PactDir:  "./pacts",
	})

	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given    []models.ProviderState
		metadata map[string]string
		content  Map
		on       func(m mocks)
	}{
		"a BasketCheckedOut message": {
			metadata: map[string]string{
				"subject": basketspb.BasketAggregateChannel,
			},
			content: Map{
				"Name": String(basketspb.BasketCheckedOutEvent),
				"Payload": Like(Map{
					"id":         String("order-id"),
					"user_id":    String("user-id"),
					"payment_id": String("payment-id"),
					"items": ArrayContaining([]interface{}{Map{
						"store_id":     String("store-id"),
						"product_id":   String("product-id"),
						"store_name":   String("store-name"),
						"product_name": String("product-name"),
						"price":        Decimal(20),
						"quantity":     Decimal(2),
					}}),
				}),
			},
			on: func(m mocks) {
				m.app.On("CreateOrder", mock.Anything, mock.AnythingOfType("commands.CreateOrder")).Return(nil)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				app: app.NewMockApp(t),
			}
			if tc.on != nil {
				tc.on(m)
			}

			handlers := integrationHandlers[ddd.Event]{
				app: m.app,
			}

			msgConsumerFn := func(contents v4.AsynchronousMessage) error {
				event := contents.Body.(*rawEvent)

				data, err := json.Marshal(event.Payload)
				if err != nil {
					return err
				}
				payload := reg.MustDeserialize(event.Name, data)

				return handlers.HandleEvent(
					context.Background(),
					ddd.NewEvent(event.Name, payload),
				)
			}
			message := pact.AddAsynchronousMessage()
			for _, given := range tc.given {
				message = message.GivenWithParameter(given)
			}
			assert.NoError(t, message.
				ExpectsToReceive(name).
				WithMetadata(tc.metadata).
				WithJSONContent(tc.content).
				AsType(&rawEvent{}).
				ConsumedBy(msgConsumerFn).
				Verify(t),
			)
		})
	}
}
