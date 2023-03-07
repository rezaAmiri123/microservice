package domain

import "github.com/rezaAmiri123/microservice/pkg/es"

type ProductV1 struct {
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

var _ es.Snapshot = (*ProductV1)(nil)

// SnapshotName implements es.Snapshot
func (ProductV1) SnapshotName() string { return "stores.ProductV1" }
