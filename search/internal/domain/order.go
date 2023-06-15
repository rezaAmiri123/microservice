package domain

import (
	"time"
)

type Order struct {
	OrderID   string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	Username  string    `bson:"username"`
	Items     []Item    `bson:"items"`
	Total     float64   `bson:"total"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
}

type Item struct {
	ProductID   string  `bson:"product_id"`
	StoreID     string  `bson:"store_id"`
	ProductName string  `bson:"product_name"`
	StoreName   string  `bson:"store_name"`
	Price       float64 `bson:"price"`
	Quantity    int     `bson:"quantity"`
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
