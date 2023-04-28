package domain

import (
	"time"
)

type Order struct {
	OrderID   string
	UserID    string
	Username  string
	Items     []Item
	Total     float64
	Status    string
	CreatedAt time.Time
}

type Item struct {
	ProductID   string
	StoreID     string
	ProductName string
	StoreName   string
	Price       float64
	Quantity    int
}

type (
	Filters struct {
		UserID     string
		After      time.Time
		Before     time.Time
		StoreIDs   []string
		ProductIDs []string
		MinTotal   float64
		MaxTotal   float64
		Status     string
	}
	SearchOrders struct {
		Filters Filters
		Next    string
		Limit   int
	}
)
