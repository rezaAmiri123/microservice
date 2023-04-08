package pg

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
)

type PaymentRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(tableName string, db postgres.DB) PaymentRepository {
	return PaymentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r PaymentRepository) Find(ctx context.Context, paymentID string) (*domain.Payment, error) {
	const query = "SELECT user_id, amount FROM %s WHERE id = $1 LIMIT 1"

	payment := &domain.Payment{
		ID: paymentID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), paymentID).Scan(
		&payment.UserID,
		&payment.Amount,
	)

	return payment, err
}

func (r PaymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	const query = "INSERT INTO %s (id, user_id, amount) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, r.table(query),
		payment.ID,
		payment.UserID,
		payment.Amount,
	)

	return err
}

func (r PaymentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
