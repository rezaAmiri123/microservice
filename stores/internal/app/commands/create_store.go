package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type (
	CreateStore struct {
		ID       string
		Name     string
		Location string
	}

	CreateStoreHandler struct {
		stores    domain.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCreateStoreHandler(stores domain.StoreRepository, publisher ddd.EventPublisher[ddd.Event], logger logger.Logger) *CreateStoreHandler {
	return &CreateStoreHandler{
		stores:    stores,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CreateStoreHandler) Handle(ctx context.Context, cmd CreateStore) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateStoreHandler.Handle")
	defer span.Finish()

	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.InitStore(cmd.Name, cmd.Location)
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
