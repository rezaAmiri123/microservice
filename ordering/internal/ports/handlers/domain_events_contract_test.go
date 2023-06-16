//go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"os"
	"path/filepath"
	"testing"
)

var pactBrokerURL string
var pactUser string
var pactPass string
var pactToken string

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/pacts", dir)

func init() {
	getEnv := func(key, fallback string) string {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
		return fallback
	}

	pactBrokerURL = getEnv("PACT_URL", "http://127.0.0.1:9292")
	pactUser = getEnv("PACT_USER", "pactuser")
	pactPass = getEnv("PACT_PASS", "pactpass")
	pactToken = getEnv("PACT_TOKEN", "")
}

func TestOrdersProducer(t *testing.T) {
	var err error

	reg := registry.New()
	err = domain.Registrations(reg)
	if err != nil {
		t.Fatal(err)
	}

	orders := domain.NewFakeOrderRepository()
	//payments := domain.NewFakePaymentRepository()
	//users := domain.NewFakeUserRepository()
	shopping := domain.NewFakeShoppingRepository()
	appLogger := applogger.NewAppLogger(applogger.Config{})

	type rawEvent struct {
		Name    string
		Payload json.RawMessage
	}

	err = orderingpb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	verifier := provider.NewVerifier()
	err = verifier.VerifyProvider(t, provider.VerifyRequest{
		PactFiles:                  []string{filepath.ToSlash(fmt.Sprintf("%s/orders-sub-pub.json", pactDir))},
		Provider:                   "orders-pub",
		ProviderVersion:            "1.0.0",
		BrokerURL:                  pactBrokerURL,
		BrokerUsername:             pactUser,
		BrokerPassword:             pactPass,
		BrokerToken:                pactToken,
		PublishVerificationResults: true,
		AfterEach: func() error {
			orders.Reset()
			shopping.Reset()
			return nil
		},
		MessageHandlers: map[string]message.Handler{
			"a OrderCanceled message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(orders, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				order := getFakeOrder()
				order.Status = domain.OrderIsPending
				orders.Reset(order)

				err = application.CancelOrder(context.Background(), commands.CancelOrder{
					ID: order.ID(),
				})
				if err != nil {
					t.Fatal(err)
				}
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}
				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]interface{}{
						"subject": subject,
					}, nil
			},
			"a OrderCompleted message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(orders, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				order := getFakeOrder()
				orders.Reset(order)
				err = application.CompleteOrder(context.Background(), commands.CompleteOrder{
					ID:        order.ID(),
					InvoiceID: order.InvoiceID,
				})
				if err != nil {
					t.Fatal(err)
				}
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}
				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]interface{}{
						"subject": subject,
					}, nil
			},
			"a OrderReadied message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(orders, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				order := getFakeOrder()
				orders.Reset(order)
				err = application.ReadyOrder(context.Background(), commands.ReadyOrder{
					ID: order.ID(),
				})
				if err != nil {
					t.Fatal(err)
				}
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}
				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]interface{}{
						"subject": subject,
					}, nil
			},
			"a OrderCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(orders, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				order := getFakeOrder()
				err := application.CreateOrder(context.Background(), commands.CreateOrder{
					ID:        order.ID(),
					UserID:    order.UserID,
					PaymentID: order.PaymentID,
					Items:     order.Items,
				})
				if err != nil {
					return nil, nil, err
				}

				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}
				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]interface{}{
						"subject": subject,
					}, nil
			},
		},
	})
}

func getFakeOrder() *domain.Order {
	order := domain.NewOrder("order-id")
	order.UserID = "user-id"
	order.PaymentID = "payment-id"
	order.InvoiceID = "invoice-id"
	order.Items = []domain.Item{{
		StoreID:     "store-id",
		StoreName:   "store-name",
		ProductID:   "product-id",
		ProductName: "product-name",
		Price:       10,
		Quantity:    2,
	}}
	return order
}
