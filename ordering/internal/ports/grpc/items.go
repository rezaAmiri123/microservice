package grpc

import (
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
)

func (s server) itemToDomain(item *orderingpb.Item) domain.Item {
	return domain.Item{
		ProductID:   item.GetProductId(),
		StoreID:     item.GetStoreId(),
		StoreName:   item.GetStoreName(),
		ProductName: item.GetProductName(),
		Price:       item.GetPrice(),
		Quantity:    int(item.GetQuantity()),
	}
}

func (s server) itemFromDomain(item domain.Item) *orderingpb.Item {
	return &orderingpb.Item{
		StoreId:     item.StoreID,
		ProductId:   item.ProductID,
		StoreName:   item.StoreName,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    int32(item.Quantity),
	}
}
