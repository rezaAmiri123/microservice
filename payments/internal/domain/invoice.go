package domain

type InvoiceStatus string

const (
	InvoiceStatusUnknown InvoiceStatus = ""
	InvoiceIsPending     InvoiceStatus = "pending"
	InvoiceIsPaid        InvoiceStatus = "paid"
	InvoiceIsCanceled    InvoiceStatus = "canceled"
)

type Invoice struct {
	ID      string
	OrderID string
	Amount  float64
	Status  InvoiceStatus
}

func (s InvoiceStatus) String() string {
	switch s {
	case InvoiceIsCanceled, InvoiceIsPaid, InvoiceIsPending:
		return string(s)
	default:
		return ""
	}
}

func ToInvoiceStatus(status string) InvoiceStatus {
	switch status {
	case InvoiceIsCanceled.String():
		return InvoiceIsCanceled
	case InvoiceIsPaid.String():
		return InvoiceIsPaid
	case InvoiceIsPending.String():
		return InvoiceIsPending
	default:
		return InvoiceStatusUnknown
	}
}
