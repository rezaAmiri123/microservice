package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type (
	RebrandStore struct {
		ID   string
		Name string
	}

	RebrandStoreHandler struct {
		stores    domain.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewRebrandStoreHandler(
	stores domain.StoreRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) RebrandStoreHandler {
	return RebrandStoreHandler{
		stores:    stores,
		publisher: publisher,
		logger:    logger,
	}
}

func (h RebrandStoreHandler) RebrandStore(ctx context.Context, cmd RebrandStore) error {

	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.Rebrand(cmd.Name)
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
