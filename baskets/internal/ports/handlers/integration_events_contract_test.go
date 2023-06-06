///go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"github.com/pact-foundation/pact-go/v2/matchers"
	v4 "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type String = matchers.String
type Map = matchers.Map

var Like = matchers.Like

func TestStoresConsumer(t *testing.T) {
	type mocks struct {
		stores   *domain.MockStoreCacheRepository
		products *domain.MockProductCacheRepository
	}

	type rawEvent struct {
		Name    string
		Payload map[string]any
	}

	reg := registry.New()
	err := storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}
	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "baskets-sub",
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
		"a StoreCreated message": {
			metadata: map[string]string{
				"subject": storespb.StoreAggregateChannel,
			},
			content: Map{
				"Name": String(storespb.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("NewStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore").Return(nil)
			},
		},
		"a StoreRebranded message": {
			metadata: map[string]string{
				"subject": storespb.StoreAggregateChannel,
			},
			content: Map{
				"Name": String(storespb.StoreRebrandedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("RebrandedStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Rename", mock.Anything, "store-id", "RebrandedStore").Return(nil)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   domain.NewMockStoreCacheRepository(t),
				products: domain.NewMockProductCacheRepository(t),
			}

			if tc.on != nil {
				tc.on(m)
			}
			handlers := integrationHandlers[ddd.Event]{m.stores, m.products}
			msgconsumerFn := func(contents v4.AsynchronousMessage) error {
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
				ConsumedBy(msgconsumerFn).
				Verify(t),
			)
		})
	}
}
