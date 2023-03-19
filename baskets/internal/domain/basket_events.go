package domain

type (
	BasketStarted struct {
		UserID string
	}

	BasketItemAdded struct {
		Item Item
	}

	BasketItemRemoved struct {
		ProductID string
		Quantity  int
	}

	BasketCanceled struct{}

	BasketCheckedOut struct {
		PaymentID string
	}
)
