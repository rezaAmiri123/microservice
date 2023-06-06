package commands

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateOrderHandler_CreateOrder(t *testing.T) {
	items := []domain.Item{{
		StoreID:     "store-id",
		StoreName:   "store-name",
		ProductID:   "product-id",
		ProductName: "product-name",
		Price:       10,
		Quantity:    2,
	}}
	type mocks struct {
		orders    *domain.MockOrderRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
		logger    *logger.MockLogger
	}

	type args struct {
		ctx context.Context
		cmd CreateOrder
	}

	tests := map[string]struct {
		args    args
		on      func(m mocks)
		wantErr bool
	}{
		"OK": {
			args: args{
				ctx: context.Background(),
				cmd: CreateOrder{
					ID:        "order-id",
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				},
			},
			on: func(m mocks) {
				m.orders.On("Load", mock.Anything, "order-id").Return(&domain.Order{
					Aggregate: es.NewAggregate("order-id", domain.OrderAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				}, nil)

				m.orders.On("Save", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(nil)
				m.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)
			},
		},
		"Publisher.Error": {
			args: args{
				ctx: context.Background(),
				cmd: CreateOrder{
					ID:        "order-id",
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				},
			},
			on: func(m mocks) {
				m.orders.On("Load", mock.Anything, "order-id").Return(&domain.Order{
					Aggregate: es.NewAggregate("order-id", domain.OrderAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				}, nil)

				m.orders.On("Save", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(nil)
				m.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(errors.New(""))
			},
			wantErr: true,
		},
		"Save.Error": {
			args: args{
				ctx: context.Background(),
				cmd: CreateOrder{
					ID:        "order-id",
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				},
			},
			on: func(m mocks) {
				m.orders.On("Load", mock.Anything, "order-id").Return(&domain.Order{
					Aggregate: es.NewAggregate("order-id", domain.OrderAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				}, nil)

				m.orders.On("Save", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(errors.New(""))
			},
			wantErr: true,
		},
		"DomainEvent.Error": {
			args: args{
				ctx: context.Background(),
				cmd: CreateOrder{
					ID: "order-id",
					//UserID:    "user-id", // do not have a user id is an error
					PaymentID: "payment-id",
					Items:     items,
				},
			},
			on: func(m mocks) {
				m.orders.On("Load", mock.Anything, "order-id").Return(&domain.Order{
					Aggregate: es.NewAggregate("order-id", domain.OrderAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				}, nil)
			},
			wantErr: true,
		},
		"Load.Error": {
			args: args{
				ctx: context.Background(),
				cmd: CreateOrder{
					ID:        "order-id",
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     items,
				},
			},
			on: func(m mocks) {
				m.orders.On("Load", mock.Anything, "order-id").Return(nil, errors.New(""))
			},
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				orders:    domain.NewMockOrderRepository(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
				logger:    logger.NewMockLogger(t),
			}

			if tc.on != nil {
				tc.on(m)
			}

			h := NewCreateOrderHandler(m.orders, m.publisher, m.logger)

			if err := h.CreateOrder(tc.args.ctx, tc.args.cmd); (err != nil) != tc.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
