package pg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/stackus/errors"
)

type OrderRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(tableName string, db postgres.DB) OrderRepository {
	return OrderRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r OrderRepository) Find(ctx context.Context, id string) (*domain.Order, error) {
	const query = "SELECT user_id, payment_id, invoice_id, shopping_id, items, status FROM %s WHERE id = $1 LIMIT 1"

	order := domain.NewOrder(id)

	var items []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), id).Scan(
		&order.UserID,
		&order.PaymentID,
		&order.InvoiceID,
		&order.ShoppingID,
		&items,
		&status,
	)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	order.Status = domain.ToOrderStatus(status)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(items, &order.Items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return order, nil
}

func (r OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	const query = "INSERT INTO %s (id, user_id, payment_id, invoice_id, shopping_id, items, status) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	items, err := json.Marshal(order.Items)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		order.ID(),
		order.UserID,
		order.PaymentID,
		order.InvoiceID,
		order.ShoppingID,
		items,
		order.Status.String(),
	)

	return errors.ErrInternalServerError.Err(err)
}

func (r OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	const query = "UPDATE %s SET items = $2, status = $3 WHERE id = $1"

	items, err := json.Marshal(order.Items)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), order.ID(), items, order.Status.String())

	return errors.ErrInternalServerError.Err(err)
}

func (r OrderRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
