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

func TestAddItemHandler_Handle(t *testing.T) {
	product := &domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	store := &domain.Store{
		ID:   "store-id",
		Name: "store-name",
	}

	type mocks struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		logger    *logger.MockLogger
		publisher *ddd.MockEventPublisher[ddd.Event]
	}

	type args struct {
		ctx context.Context
		add AddItem
	}

	tests := map[string]struct {
		args    args
		on      func(m mocks)
		wantErr bool
	}{
		"Success": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					BasketID:  "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     make(map[string]domain.Item),
					Status:    domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(store, nil)
				f.baskets.On("UpdateItems", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
			},
		},
		"NoBasket": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					BasketID:  "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(nil, fmt.Errorf("no basket"))
			},
			wantErr: true,
		},
		"NoProduct": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					BasketID:  "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     make(map[string]domain.Item),
					Status:    domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(nil, fmt.Errorf("no product"))
			},
			wantErr: true,
		},
		"NoStore": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					BasketID:  "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     make(map[string]domain.Item),
					Status:    domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(nil, fmt.Errorf("no store"))
			},
			wantErr: true,
		},
		"SaveFailed": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					BasketID:  "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
					UserID:    "user-id",
					PaymentID: "payment-id",
					Items:     make(map[string]domain.Item),
					Status:    domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(store, nil)
				f.baskets.On("UpdateItems", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(fmt.Errorf("save failed"))
			},
			wantErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				baskets:   domain.NewMockBasketRepository(t),
				stores:    domain.NewMockStoreRepository(t),
				products:  domain.NewMockProductRepository(t),
				logger:    logger.NewMockLogger(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
			}

			a := NewAddItemHandler(m.baskets, m.stores, m.products, m.publisher, m.logger)

			if tc.on != nil {
				tc.on(m)
			}

			if err := a.AddItem(tc.args.ctx, tc.args.add); (err != nil) != tc.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
