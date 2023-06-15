package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FakeStoreServer struct {
	stores   map[string]*storespb.Store
	products map[string]*storespb.Product
	storespb.UnimplementedStoresServiceServer
}

func NewFakeStoreServer() *FakeStoreServer {
	return &FakeStoreServer{
		stores:   make(map[string]*storespb.Store),
		products: make(map[string]*storespb.Product),
	}
}

func (r *FakeStoreServer) GetStore(ctx context.Context, req *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
	if store, exists := r.stores[req.GetId()]; exists {
		return &storespb.GetStoreResponse{
			Store: store,
		}, nil
	}
	return nil, status.Error(codes.NotFound, "store not exists")
}
func (r *FakeStoreServer) GetProduct(ctx context.Context, req *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
	if product, exists := r.products[req.GetId()]; exists {
		return &storespb.GetProductResponse{
			Product: product,
		}, nil
	}

	return nil, status.Error(codes.NotFound, "product not exists")
}

func (r *FakeStoreServer) AddStores(stores ...*storespb.Store) {
	for _, store := range stores {
		r.stores[store.GetId()] = store
	}
}
func (r *FakeStoreServer) AddProducts(products ...*storespb.Product) {
	for _, product := range products {
		r.products[product.GetId()] = product
	}
}

func (r *FakeStoreServer) Reset() {
	r.stores = make(map[string]*storespb.Store)
	r.products = make(map[string]*storespb.Product)
}
