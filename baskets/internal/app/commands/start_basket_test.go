package commands

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestStartBasketHandler_Handle(t *testing.T) {
	type mocks struct {
		baskets *domain.MockBasketRepository
		//stores    *domain.MockStoreRepository
		//products  *domain.MockProductRepository
		logger    *logger.MockLogger
		publisher *ddd.MockEventPublisher[ddd.Event]
	}

	type args struct {
		ctx   context.Context
		start StartBasket
	}

	tests := map[string]struct {
		args    args
		on      func(m mocks)
		wantErr bool
	}{
		"Success": {
			args: args{
				ctx: context.Background(),
				start: StartBasket{
					ID:     "basket-id",
					UserID: "user-id",
				},
			},
			on: func(m mocks) {
				m.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "",
					Items:     make(map[string]domain.Item),
				}, nil)
				m.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
				m.publisher.On("Publish", context.Background(), mock.AnythingOfType("ddd.event")).Return(nil)
			},
		},
		"NoBasket": {
			args: args{
				ctx: context.Background(),
				start: StartBasket{
					ID:     "basket-id",
					UserID: "user-id",
				},
			},
			on: func(m mocks) {
				m.baskets.On("Load", context.Background(), "basket-id").Return(nil, fmt.Errorf("no basket"))
			},
			wantErr: true,
		},
		"SaveFailed": {
			args: args{
				ctx: context.Background(),
				start: StartBasket{
					ID:     "basket-id",
					UserID: "user-id",
				},
			},
			on: func(m mocks) {
				m.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "",
					Items:     make(map[string]domain.Item),
				}, nil)
				m.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(fmt.Errorf("save failed"))
			},
			wantErr: true,
		},
		"PublishFailed": {
			args: args{
				ctx: context.Background(),
				start: StartBasket{
					ID:     "basket-id",
					UserID: "user-id",
				},
			},
			on: func(m mocks) {
				m.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "",
					Items:     make(map[string]domain.Item),
				}, nil)
				m.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
				m.publisher.On("Publish", context.Background(), mock.AnythingOfType("ddd.event")).Return(fmt.Errorf("publish failed"))
			},
			wantErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				baskets:   domain.NewMockBasketRepository(t),
				logger:    logger.NewMockLogger(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
			}

			a := NewStartBasketHandler(m.baskets, m.publisher, m.logger)

			if tc.on != nil {
				tc.on(m)
			}

			if err := a.StartBasket(tc.args.ctx, tc.args.start); (err != nil) != tc.wantErr {
				t.Errorf("StartBasket error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
