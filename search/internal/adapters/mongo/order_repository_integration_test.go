// /go:build integration || database
package mongo

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/mongodb"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

const (
	testDBName         = "test"
	testCollectionName = "users"
)

type orderSuite struct {
	container testcontainers.Container
	db        *mongo.Client
	//mock      *app.MockUserRepository
	repo OrderRepository
	suite.Suite
}

func TestOrderRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("short mode: skipping")
	}
	suite.Run(t, &orderSuite{})
}

func (s *orderSuite) SetupSuite() {
	var err error

	ctx := context.Background()
	s.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo:latest",
			Hostname:     "mongodb",
			ExposedPorts: []string{"27017/tcp"},
			Env: map[string]string{
				"MONGO_INITDB_ROOT_USERNAME": "admin",
				"MONGO_INITDB_ROOT_PASSWORD": "admin",
				"MONGODB_DATABASE":           "test",
			},
			WaitingFor: wait.ForLog("Waiting for connections").WithStartupTimeout(5 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	endpoint, err := s.container.Endpoint(ctx, "")
	if err != nil {
		s.T().Fatal(err)
	}
	mongoDBConn, err := mongodb.NewMongoDBConn(context.Background(), &mongodb.Config{
		URI:      fmt.Sprintf("mongodb://%s", endpoint),
		User:     "admin",
		Password: "admin",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = mongoDBConn
}
func (s *orderSuite) TearDownSuite() {
	err := s.db.Disconnect(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
	err = s.container.Terminate(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
}
func (s *orderSuite) SetupTest() {
	//s.mock = app.NewMockUserRepository(s.T())
	s.repo = NewOrderRepository(testDBName, testCollectionName, s.db)
}

func (s *orderSuite) TearDownTest() {
	err := s.getCollection().Drop(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
}
func (s *orderSuite) getOrder() *domain.Order {
	return &domain.Order{
		OrderID:   "order-id",
		UserID:    "user-id",
		Username:  "username",
		Status:    "New",
		Total:     10,
		CreatedAt: time.Now(),
		Items: []domain.Item{{
			ProductID:   "product-id",
			ProductName: "product-name",
			StoreID:     "store-id",
			StoreName:   "store-name",
			Quantity:    2,
			Price:       10.10,
		}},
	}
}
func (s *orderSuite) TestOrderRepository_Add() {
	err := s.repo.Add(context.Background(), s.getOrder())
	s.NoError(err)
	order := &domain.Order{}

	err = s.getCollection().FindOne(context.Background(), bson.M{"_id": "order-id"}).Decode(order)
	s.NoError(err)
	s.Equal("username", order.Username)
	s.Equal("user-id", order.UserID)
}

func (s *orderSuite) TestOrderRepository_Get() {
	order := s.getOrder()

	_, err := s.getCollection().InsertOne(context.Background(), order, &options.InsertOneOptions{})
	s.NoError(err)

	newOrder, err := s.repo.Get(context.Background(), order.OrderID)
	s.NoError(err)
	s.NotNil(newOrder)
	s.Equal(order.OrderID, newOrder.OrderID)
	s.Equal(order.Username, newOrder.Username)
}

func (s *orderSuite) TestOrderRepository_UpdateStatus() {
	const newStatus = "NewStatus"
	order := s.getOrder()

	_, err := s.getCollection().InsertOne(context.Background(), order, &options.InsertOneOptions{})
	s.NoError(err)

	err = s.repo.UpdateStatus(context.Background(), order.OrderID, newStatus)
	s.NoError(err)
	newOrder := &domain.Order{}
	err = s.getCollection().FindOne(context.Background(), bson.M{"_id": order.OrderID}).Decode(newOrder)
	s.NoError(err)
	s.Equal(newOrder.Username, order.Username)
	s.Equal(newOrder.OrderID, order.OrderID)
	s.Equal(newStatus, newOrder.Status)
}

func (s *orderSuite) TestOrderRepository_Search() {

	for i := 0; i < 4; i++ {
		order := s.getOrder()
		order.OrderID = fmt.Sprintf("order-id-%d", i)
		_, err := s.getCollection().InsertOne(context.Background(), order, &options.InsertOneOptions{})
		s.NoError(err)
	}

	orders, err := s.repo.Search(context.Background(), domain.SearchOrders{
		Filters: domain.Filters{
			Status: "New",
			UserID: "user-id",
		},
		Limit: 2,
	})
	s.NoError(err)
	for _, newOrder := range orders {
		fmt.Println(*newOrder)
	}
}

//	func (s *orderSuite) TestUserCacheRepository_FindFromFallback() {
//		s.mock.On("Find", mock.Anything, "user-id").Return(&models.User{
//			ID:       "user-id",
//			Username: "username",
//		}, nil)
//
//		newUser, err := s.repo.Find(context.Background(), "user-id")
//		s.NoError(err)
//		s.NotNil(newUser)
//		s.Equal("user-id", newUser.ID)
//		s.Equal("username", newUser.Username)
//	}
func (s *orderSuite) getCollection() *mongo.Collection {
	return s.db.Database(testDBName).Collection(testCollectionName)
}
