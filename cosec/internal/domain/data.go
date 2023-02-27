package domain

type CreateOrderData struct {
	//OrderID    string
	UserID string
	//PaymentID  string
	//ShoppingID string
	Items []Item
	//Total      float64
}

type Item struct {
	//ProductID string
	//StoreID   string
	//Price     float64
	//Quantity  int
}
