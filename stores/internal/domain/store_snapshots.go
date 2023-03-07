package domain

import "github.com/rezaAmiri123/microservice/pkg/es"

type StoreV1 struct {
	Name          string
	Location      string
	Participating bool
}

var _ es.Snapshot = (*StoreV1)(nil)

// SnapshotName implements es.Snapshot
func (StoreV1) SnapshotName() string { return "stores.StoreV1" }
