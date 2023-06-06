package domain

import "context"

type FakePaymentRepository struct {
}

func NewFakePaymentRepository() *FakePaymentRepository {
	return &FakePaymentRepository{}
}

var _ PaymentRepository = (*FakePaymentRepository)(nil)

func (r *FakePaymentRepository) Confirm(ctx context.Context, paymentID string) error {
	return nil
}
