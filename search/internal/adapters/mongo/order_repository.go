package mongo

import (
	"context"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/stackus/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	dbName    string
	tableName string
	db        *mongo.Client
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(dbName, tableName string, db *mongo.Client) OrderRepository {
	return OrderRepository{
		dbName:    dbName,
		tableName: tableName,
		db:        db,
	}
}

func (r OrderRepository) Add(ctx context.Context, order *domain.Order) error {
	_, err := r.getCollection().InsertOne(ctx, order, &options.InsertOneOptions{})

	return err
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	ops := options.FindOneAndUpdate()
	//ops.SetReturnDocument(options.After)
	ops.SetUpsert(true)

	return r.getCollection().FindOneAndUpdate(ctx, bson.M{"_id": orderID}, bson.M{"$set": bson.M{
		"status": status,
	}}, ops).Err()
}

func (r OrderRepository) getFilters(filters domain.Filters) []bson.M {
	query := []bson.M{}
	if filters.Status != "" {
		query = append(query, bson.M{"status": filters.Status})
	}
	if filters.UserID != "" {
		query = append(query, bson.M{"user_id": filters.UserID})
	}
	return query
}

func (r OrderRepository) Search(ctx context.Context, search domain.SearchOrders) ([]*domain.Order, error) {
	query := r.getFilters(search.Filters)
	limit := int64(search.Limit)
	//skip := int64(search.Next)
	opts := &options.FindOptions{
		Limit: &limit,
		//Skip: skip
	}
	cursor, err := r.getCollection().Find(ctx, bson.M{"$and": query}, opts)
	if err != nil {
		return nil, err
	}
	orders := []*domain.Order{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, errors.Wrap(err, "Search")
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r OrderRepository) Get(ctx context.Context, orderID string) (*domain.Order, error) {
	order := &domain.Order{}
	err := r.getCollection().FindOne(ctx, bson.M{"_id": orderID}).Decode(order)
	return order, err
}

func (r OrderRepository) getCollection() *mongo.Collection {
	return r.db.Database(r.dbName).Collection(r.tableName)
}
