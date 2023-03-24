package pg

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/stackus/errors"
)

type InvoiceRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.InvoiceRepository = (*InvoiceRepository)(nil)

func NewInvoiceRepository(tableName string, db postgres.DB) InvoiceRepository {
	return InvoiceRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r InvoiceRepository) Find(ctx context.Context, invoiceID string) (*domain.Invoice, error) {
	const query = "SELECT order_id, amount, status FROM %s WHERE id = $1 LIMIT 1"

	invoice := &domain.Invoice{
		ID: invoiceID,
	}
	var status string
	err := r.db.QueryRowContext(ctx, r.table(query), invoiceID).Scan(&invoice.OrderID, &invoice.Amount, &status)
	if err != nil {
		return nil, errors.Wrap(err, "scanning invoice")
	}

	invoice.Status = domain.ToInvoiceStatus(status)

	return invoice, nil
}

func (r InvoiceRepository) Save(ctx context.Context, invoice *domain.Invoice) error {
	const query = "INSERT INTO %s (id, order_id, amount, status) VALUES ($1, $2, $3, $4)"
	fmt.Println(*invoice)
	_, err := r.db.ExecContext(ctx, r.table(query), invoice.ID, invoice.OrderID, invoice.Amount, invoice.Status.String())
	fmt.Println(err)
	return err
}

func (r InvoiceRepository) Update(ctx context.Context, invoice *domain.Invoice) error {
	const query = "UPDATE %s SET amount = $2, status = $3 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), invoice.ID, invoice.Amount, invoice.Status.String())

	return err
}

func (r InvoiceRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
