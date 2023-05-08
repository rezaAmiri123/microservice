package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type (
	DisableParticipation struct {
		ID string
	}

	DisableParticipationHandler struct {
		stores    domain.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewDisableParticipationHandler(
	stores domain.StoreRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) DisableParticipationHandler {
	return DisableParticipationHandler{
		stores:    stores,
		publisher: publisher,
		logger:    logger,
	}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {

	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.DisableParticipation()
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
