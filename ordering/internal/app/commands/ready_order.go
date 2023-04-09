package commands

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	ReadyOrder struct {
		ID string
	}

	ReadyOrderHandler struct {
		orders    domain.OrderRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewReadyOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *ReadyOrderHandler {
	fmt.Println(publisher)
	return &ReadyOrderHandler{
		orders:    orders,
		publisher: publisher,
		logger:    logger,
	}
}

func (h ReadyOrderHandler) Handle(ctx context.Context, cmd ReadyOrder) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ReadyOrderHandler.Handle")
	defer span.Finish()

	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Ready()
	if err != nil {
		return errors.Wrap(err, "ready order command")
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "order ready")
	}

	err = h.publisher.Publish(ctx, event)

	return err
}