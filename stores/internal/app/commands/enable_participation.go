package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type (
	EnableParticipation struct {
		ID string
	}

	EnableParticipationHandler struct {
		stores    domain.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewEnableParticipationHandler(
	stores domain.StoreRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) EnableParticipationHandler {
	return EnableParticipationHandler{
		stores:    stores,
		publisher: publisher,
		logger:    logger,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {

	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.EnableParticipation()
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
