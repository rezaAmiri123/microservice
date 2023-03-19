package domain

import "github.com/rezaAmiri123/microservice/pkg/es"

type BasketV1 struct {
	UserID    string
	PaymentID string
	Items     map[string]Item
	Status    BasketStatus
}

var _ es.Snapshot = (*BasketV1)(nil)

// SnapshotName implements es.Snapshot
func (BasketV1) SnapshotName() string { return "baskets.BasketV1" }
