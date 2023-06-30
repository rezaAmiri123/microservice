package sec

import (
	"context"
	"github.com/stackus/errors"
)

type SagaKey struct {
	ID string
	Name string
}
type FakeSagaStoreRepository struct {
	sagaStore map[SagaKey]*SagaContext[[]byte]
}

func NewFakeSagaStoreRepository() *FakeSagaStoreRepository {
	return &FakeSagaStoreRepository{
		sagaStore: make(map[SagaKey]*SagaContext[[]byte]),
	}
}

var _ SagaStore = (*FakeSagaStoreRepository)(nil)

func (r *FakeSagaStoreRepository)Load(ctx context.Context, sagaName, sagaID string) (*SagaContext[[]byte], error)  {
	key:= SagaKey{ID: sagaID, Name: sagaName}
	if sagaCtx , exists := r.sagaStore[key];exists{
		return sagaCtx,nil
	}
	return nil, errors.ErrNotFound.Msgf("saga context with id: %v, name: %v not found", sagaID, sagaName)
}

func (r *FakeSagaStoreRepository)  Save(ctx context.Context, sagaName string, sagaCtx *SagaContext[[]byte]) error{
	key:= SagaKey{ID: sagaCtx.ID, Name: sagaName}
	r.sagaStore[key]=sagaCtx
	return nil
}
func (r *FakeSagaStoreRepository)  Reset(name string, sagaCtx *SagaContext[[]byte]){
	r.sagaStore = make(map[SagaKey]*SagaContext[[]byte])
	_ = r.Save(context.Background(), name, sagaCtx)
}
