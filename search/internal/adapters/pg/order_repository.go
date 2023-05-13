package pg

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/stackus/errors"
	"strings"
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

func (r OrderRepository) Add(ctx context.Context, order *domain.Order) error {
	const query = `INSERT INTO %s (
order_id, user_id, username,
items, status, product_ids, store_ids,
 total) VALUES (
$1, $2, $3,
$4, $5, $6, $7,
$8)`

	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	productIDs := make(IDArray, len(order.Items))
	storeMap := make(map[string]struct{})
	for i, item := range order.Items {
		productIDs[i] = item.ProductID
		storeMap[item.StoreID] = struct{}{}
	}
	storeIDs := make(IDArray, 0, len(storeMap))
	for storeID, _ := range storeMap {
		storeIDs = append(storeIDs, storeID)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		order.OrderID, order.UserID, order.Username,
		items, order.Status, productIDs, storeIDs,
		order.Total,
	)
	return err
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	const query = `UPDATE %s SET status = $2 WHERE order_id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), orderID, status)
	return err
}
func (r OrderRepository) getFilters(search domain.SearchOrders) (query string, filterArgs []any) {
	filters := search.Filters
	var filterNum = 1

	if filters.Status != "" {
		query += fmt.Sprintf("status = $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.Status)
	}

	if filters.UserID != "" {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("user_id = $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.UserID)
	}

	if len(filters.StoreIDs) > 0 {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("store_ids IN ($%d) ", filterNum)
		filterNum++
		storeIDs := make(IDArray, 0)
		for _, id := range filters.StoreIDs {
			storeIDs = append(storeIDs, id)
		}
		filterArgs = append(filterArgs, storeIDs)
	}

	if len(filters.ProductIDs) > 0 {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("product_ids IN ($%d) ", filterNum)
		filterNum++
		productIDs := make(IDArray, 0)
		for _, id := range filters.ProductIDs {
			productIDs = append(productIDs, id)
		}
		filterArgs = append(filterArgs, productIDs)
	}

	if !filters.After.IsZero() {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("created_at > $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.After)
	}

	if !filters.Before.IsZero() {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("created_at < $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.Before)
	}

	if filters.MinTotal != 0 {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("total > $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.MinTotal)
	}

	if filters.MaxTotal != 0 {
		if filterNum > 1 {
			query += fmt.Sprintf("AND ")
		}
		query += fmt.Sprintf("total < $%d ", filterNum)
		filterNum++
		filterArgs = append(filterArgs, filters.MaxTotal)
	}

	if search.Limit == 0 {
		search.Limit = 10
	}
	query += fmt.Sprintf("Limit $%d ", filterNum)
	filterArgs = append(filterArgs, search.Limit)
	
	return
}
func (r OrderRepository) Search(ctx context.Context, search domain.SearchOrders) (orders []*domain.Order, err error) {
	var query = `SELECT order_id, user_id, username, items, status, created_at, total FROM %s `
	filterQuery, filterArgs := r.getFilters(search)
	if filterQuery != "" {
		query = query + fmt.Sprintf("WHERE %s", filterQuery)
	}

	var rows *sql.Rows
	rows, err = r.db.QueryContext(ctx, r.table(query), filterArgs...)
	if err != nil {
		return nil, errors.Wrap(err, "querying orders")
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing orders rows")
		}
	}(rows)

	for rows.Next() {
		order := &domain.Order{}
		var itemData []byte

		err := rows.Scan(
			&order.OrderID,
			&order.UserID,
			&order.Username,
			&itemData,
			&order.Status,
			&order.CreatedAt,
			&order.Total,
		)
		if err != nil {
			return nil, err
		}

		var items []domain.Item
		err = json.Unmarshal(itemData, &items)
		if err != nil {
			return nil, err
		}
		order.Items = items
		orders = append(orders, order)
	}
	// TODO not implemented
	return orders, nil
}

func (r OrderRepository) Get(ctx context.Context, orderID string) (*domain.Order, error) {
	const query = `SELECT user_id, username, items, status, created_at FROM %s WHERE order_id = $1`

	order := &domain.Order{
		OrderID: orderID,
	}
	var itemData []byte
	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(
		&order.UserID,
		&order.Username,
		&itemData,
		&order.Status,
		&order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	var items []domain.Item
	err = json.Unmarshal(itemData, &items)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return order, err
}

func (r OrderRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

type IDArray []string

func (a IDArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	if len(a) == 0 {
		return "{}", nil
	}
	// unsafe way to do this; assumption is all ids are UUIDs
	return fmt.Sprintf("{%s}", strings.Join(a, ",")), nil
}

func (a *IDArray) Scan(src any) error {
	var sep = []byte(",")

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.ErrInvalidArgument.Msgf("IDArray: unsupported type: %T", src)
	}

	ids := make([]string, bytes.Count(data, sep))
	for i, id := range bytes.Split(bytes.Trim(data, "{}"), sep) {
		ids[i] = string(id)
	}
	*a = ids

	return nil
}
