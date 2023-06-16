//go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/version"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/rezaAmiri123/microservice/stores/internal/app"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
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

	pactBrokerURL = getEnv("PACT_URL", "http://pact:9292")
	pactUser = getEnv("PACT_USER", "pactuser")
	pactPass = getEnv("PACT_PASS", "pactpass")
	pactToken = getEnv("PACT_TOKEN", "")
}

func TestProvider(t *testing.T) {
	version.CheckVersion()

	var err error

	// init registry
	reg := registry.New()
	err = domain.Registrations(reg)
	if err != nil {
		t.Fatal(err)
	}
	// init repos
	stores := domain.NewFakeStoreRepository()
	products := domain.NewFakeProductRepository()
	mall := domain.NewFakeMallRepository()
	catalog := domain.NewFakeCatalogRepository()
	appLogger := applogger.NewAppLogger(applogger.Config{})

	type rawEvent struct {
		Name    string
		Payload json.RawMessage
	}

	err = storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}
	verifier := provider.NewVerifier()
	err = verifier.VerifyProvider(t, provider.VerifyRequest{
		PactFiles:                  []string{filepath.ToSlash(fmt.Sprintf("%s/stores-sub-pub.json", pactDir))},
		Provider:                   "stores-pub",
		ProviderVersion:            "1.0.0",
		BrokerURL:                  pactBrokerURL,
		BrokerUsername:             pactUser,
		BrokerPassword:             pactPass,
		BrokerToken:                pactToken,
		PublishVerificationResults: true,
		AfterEach: func() error {
			stores.Reset()
			products.Reset()
			return nil
		},
		MessageHandlers: map[string]message.Handler{
			"a StoreCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(stores, products, catalog, mall, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				// Act
				err := application.CreateStore(context.Background(), commands.CreateStore{
					ID:       "store-id",
					Name:     "NewStore",
					Location: "NewLocation",
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
					}, map[string]any{
						"subject": subject,
					}, nil
			},
			"a StoreRebranded message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				application := app.New(stores, products, catalog, mall, dispatcher, appLogger)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				store := domain.NewStore("store-id")
				store.Name = "NewStore"
				store.Location = "NewLocation"
				stores.Reset(store)

				err := application.RebrandStore(context.Background(), commands.RebrandStore{
					ID:   "store-id",
					Name: "RebrandedStore",
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
					}, map[string]any{
						"subject": subject,
					}, nil
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
}
