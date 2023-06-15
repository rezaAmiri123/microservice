//go:build contract

package rest

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/baskets/internal/ports/grpc"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/stretchr/testify/assert"
	grpcstd "google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var pactBrokerURL string
var pactUser string
var pactPass string
var pactToken string

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
	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/pacts", dir)

	var err error

	// init registry
	reg := registry.New()
	err = domain.Registrations(reg)
	if err != nil {
		t.Fatal(err)
	}
	// init repos
	baskets := domain.NewFakeBasketRepository()
	stores := domain.NewFakeStoreCacheRepository()
	products := domain.NewFakeProductCacheRepository()
	dispatcher := ddd.NewEventDispatcher[ddd.Event]()
	appLogger := applogger.NewAppLogger(applogger.Config{})

	webAddress := fmt.Sprintf(":9091")
	grpcAddress := fmt.Sprintf(":9095")

	application := app.New(baskets, stores, products, dispatcher, appLogger)
	grpcServer := grpcstd.NewServer()

	mux := chi.NewMux()

	err = grpc.RegisterServer(application, grpcServer, appLogger)
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterGateway(context.Background(), mux, grpcAddress)
	if err != nil {
		t.Fatal(err)
	}

	// start up the GRPC API; we proxy the REST api through the GRPC API
	{
		listener, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			if err = grpcServer.Serve(listener); err != nil && err != grpcstd.ErrServerStopped {
				t.Error(err)
				return
			}
			defer func() {
				grpcServer.GracefulStop()
			}()
		}()

	}
	// start up the REST API
	{
		webServer := &http.Server{
			Addr:    webAddress,
			Handler: mux,
		}
		go func() {
			if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				t.Error(err)
				return
			}
		}()
	}

	verifier := provider.NewVerifier()
	assert.NoError(t, verifier.VerifyProvider(t, provider.VerifyRequest{
		PactFiles: []string{
			filepath.ToSlash(fmt.Sprintf("%s/gateway.json", pactDir)),
		},
		Provider:                   "baskets-api",
		ProviderBaseURL:            fmt.Sprintf("http://%s", webAddress),
		ProviderVersion:            "1.0.0",
		BrokerURL:                  pactBrokerURL,
		BrokerToken:                pactToken,
		BrokerUsername:             pactUser,
		BrokerPassword:             pactPass,
		PublishVerificationResults: true,
		AfterEach: func() error {
			baskets.Reset()
			products.Reset()
			stores.Reset()
			return nil
		},
		StateHandlers: map[string]models.StateHandler{
			"a basket exists": func(setup bool, state models.ProviderState) (models.ProviderStateResponse, error) {
				basket := domain.NewBasket("basket-id")
				if v, exists := state.Parameters["id"]; exists {
					basket = domain.NewBasket(v.(string))
				}
				basket.Items = map[string]domain.Item{}
				basket.UserID = "user-id"
				if v, exists := state.Parameters["userId"]; exists {
					basket.UserID = v.(string)
				}
				basket.Status = domain.BasketIsOpen
				if v, exists := state.Parameters["status"]; exists && domain.BasketStatus(v.(string)).String() != "" {
					basket.Status = domain.BasketStatus(v.(string))
				}
				baskets.Reset(basket)
				return nil, nil
			},
		},
	}))
}
