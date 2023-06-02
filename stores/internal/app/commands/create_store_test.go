package commands

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestNewCreateStoreHandler(t *testing.T) {
	stores := domain.NewMockStoreRepository(t)
	publishers := ddd.NewMockEventPublisher[ddd.Event](t)
	logs := logger.NewMockLogger(t)

	handler := NewCreateStoreHandler(stores, publishers, logs)
	assert.True(t, reflect.DeepEqual(stores, handler.stores))
	assert.True(t, reflect.DeepEqual(publishers, handler.publisher))
	assert.True(t, reflect.DeepEqual(logs, handler.logger))

}
func TestCreateStoreHandler_CreateStore(t *testing.T) {
	const (
		testStoreID       = "store-id"
		testStoreName     = "store-name"
		testStoreLocation = "store-location"
	)

	type args struct {
		ctx         context.Context
		CreateStore CreateStore
	}
	tests := map[string]struct {
		args    args
		on      func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger)
		wantErr bool
	}{
		"OK": {
			args: args{
				ctx: context.Background(),
				CreateStore: CreateStore{
					ID:       testStoreID,
					Name:     testStoreName,
					Location: testStoreLocation,
				},
			},
			on: func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger) {
				stores.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Store{
					Aggregate: es.NewAggregate(testStoreID, domain.StoreAggregate),
					Name:      testStoreName,
					Location:  testStoreLocation,
				}, nil)

				stores.On("Save", mock.Anything, mock.AnythingOfType("*domain.Store")).Return(nil)
				publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)
			},
		},
		"Store.LoadError": {
			args: args{
				ctx: context.Background(),
				CreateStore: CreateStore{
					ID:       testStoreID,
					Name:     testStoreName,
					Location: testStoreLocation,
				},
			},
			on: func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger) {
				stores.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
		"Store.InitStoreError": {
			args: args{
				ctx: context.Background(),
				CreateStore: CreateStore{
					ID: testStoreID,
					//Name:     testStoreName,
					Location: testStoreLocation,
				},
			},
			on: func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger) {
				stores.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Store{
					Aggregate: es.NewAggregate(testStoreID, domain.StoreAggregate),
					Name:      testStoreName,
					Location:  testStoreLocation,
				}, nil)
			},
			wantErr: true,
		},
		"Store.SaveError": {
			args: args{
				ctx: context.Background(),
				CreateStore: CreateStore{
					ID:       testStoreID,
					Name:     testStoreName,
					Location: testStoreLocation,
				},
			},
			on: func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger) {
				stores.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Store{
					Aggregate: es.NewAggregate(testStoreID, domain.StoreAggregate),
					Name:      testStoreName,
					Location:  testStoreLocation,
				}, nil)

				stores.On("Save", mock.Anything, mock.AnythingOfType("*domain.Store")).Return(errors.New("save error"))
			},
			wantErr: true,
		},
		"Store.PublishError": {
			args: args{
				ctx: context.Background(),
				CreateStore: CreateStore{
					ID:       testStoreID,
					Name:     testStoreName,
					Location: testStoreLocation,
				},
			},
			on: func(stores *domain.MockStoreRepository, publisher *ddd.MockEventPublisher[ddd.Event], log *logger.MockLogger) {
				stores.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Store{
					Aggregate: es.NewAggregate(testStoreID, domain.StoreAggregate),
					Name:      testStoreName,
					Location:  testStoreLocation,
				}, nil)

				stores.On("Save", mock.Anything, mock.AnythingOfType("*domain.Store")).Return(nil)
				publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			stores := domain.NewMockStoreRepository(t)
			publishers := ddd.NewMockEventPublisher[ddd.Event](t)
			logs := logger.NewMockLogger(t)

			h := CreateStoreHandler{
				stores:    stores,
				publisher: publishers,
				logger:    logs,
			}

			if tt.on != nil {
				tt.on(stores, publishers, logs)
			}

			err := h.CreateStore(tt.args.ctx, tt.args.CreateStore)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
