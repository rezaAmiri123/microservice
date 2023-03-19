package pg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
)

type BasketRepository struct {
	tableName string
	db        postgres.DB
}

var _ domain.BasketRepository = (*BasketRepository)(nil)

func NewBasketRepository(tableName string, db postgres.DB) BasketRepository {
	return BasketRepository{
		tableName: tableName,
		db:        db,
	}
}
func (r BasketRepository) Load(ctx context.Context, basketID string) (*domain.Basket, error) {
	const query = "SELECT user_id, payment_id, status, items FROM %s WHERE id = $1 LIMIT 1"

	basket := domain.NewBasket(basketID)

	var status string
	var itemsData string

	err := r.db.QueryRowContext(ctx, r.table(query), basketID).Scan(
		&basket.UserID,
		&basket.PaymentID,
		&status,
		&itemsData,
	)
	if err != nil {
		return nil, err
	}

	basket.Status = domain.ToBasketStatus(status)
	if itemsData != "" {
		err = json.Unmarshal([]byte(itemsData), &basket.Items)
		if err != nil {
			return nil, err
		}
	}

	return basket, err
}

func (r BasketRepository) Save(ctx context.Context, basket *domain.Basket) error {

	const query = "INSERT INTO %s (id, user_id, payment_id, status, items) VALUES ($1, $2, $3, $4, $5)"
	items, err := json.Marshal(basket.Items)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		basket.ID(),
		basket.UserID,
		basket.PaymentID,
		basket.Status.String(),
		items,
	)

	return err
}
func (r BasketRepository) UpdateItems(ctx context.Context, basket *domain.Basket) error {
	const query = `UPDATE %s SET items = $2 WHERE id = $1`

	items, err := json.Marshal(basket.Items)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, r.table(query), basket.ID(), items)

	return err
}

func (r BasketRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
