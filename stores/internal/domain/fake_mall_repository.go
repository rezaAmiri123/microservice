package domain

import (
	"context"
	"fmt"
	"github.com/stackus/errors"
)

type FakeMallRepository struct {
	stores map[string]*MallStore
}

var _ MallRepository = (*FakeMallRepository)(nil)

func NewFakeMallRepository() *FakeMallRepository {
	return &FakeMallRepository{
		stores: map[string]*MallStore{},
	}
}

func (r *FakeMallRepository) AddStore(ctx context.Context, storeID, name, location string) error {
	r.stores[storeID] = &MallStore{
		ID:       storeID,
		Name:     name,
		Location: location,
	}
	return nil
}

func (r *FakeMallRepository) SetStoreParticipation(ctx context.Context, storeID string, participating bool) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) RenameStore(ctx context.Context, storeID, name string) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) Find(ctx context.Context, storeID string) (*MallStore, error) {
	if store, exists := r.stores[storeID]; exists {
		return store, nil
	}
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$4444")
	fmt.Println("Find: ", r.stores)
	return nil, errors.ErrNotFound.Msgf("store with id: `%s` does not exists", storeID)
}

func (r *FakeMallRepository) All(ctx context.Context) ([]*MallStore, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) AllParticipating(ctx context.Context) ([]*MallStore, error) {
	// TODO implement me
	panic("implement me")
}
func (r *FakeMallRepository) Reset(mall ...*MallStore) {
	r.stores = make(map[string]*MallStore)
	for _, store := range mall {
		r.stores[store.ID] = store
	}
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$4444")
	fmt.Println("Reset: ", r.stores)
}
