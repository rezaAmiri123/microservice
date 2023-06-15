// /go:build integration || database
package mongo

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/models"
	"github.com/rezaAmiri123/microservice/pkg/db/mongodb"
	"github.com/stretchr/testify/mock"
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

type userCacheSuite struct {
	container testcontainers.Container
	db        *mongo.Client
	mock      *app.MockUserRepository
	repo      UserCacheRepository
	suite.Suite
}

func TestUserCacheRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("short mode: skipping")
	}
	suite.Run(t, &userCacheSuite{})
}

func (s *userCacheSuite) SetupSuite() {
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
			//Mounts: []testcontainers.ContainerMount{
			//	testcontainers.BindMount(initDir, "/docker-entrypoint-initdb.d"),
			//}, "Waiting for connections"
			//WaitingFor: wait.ForSQL("27017/tcp", "pgx", func(host string, port nat.Port) string {
			//	return fmt.Sprintf(dbUrl, host, port.Port())
			//}).WithStartupTimeout(5 * time.Second),
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
func (s *userCacheSuite) TearDownSuite() {
	err := s.db.Disconnect(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
	err = s.container.Terminate(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
}
func (s *userCacheSuite) SetupTest() {
	s.mock = app.NewMockUserRepository(s.T())
	s.repo = NewUserCacheRepository("test", "users", s.db, s.mock)
}

func (s *userCacheSuite) TearDownTest() {
	err := s.getCollection().Drop(context.Background())
	if err != nil {
		s.T().Fatal(err)
	}
}
func (s *userCacheSuite) TestUserCacheRepository_Add() {
	err := s.repo.Add(context.Background(), "user-id", "username")
	s.NoError(err)
	user := &models.User{}

	err = s.getCollection().FindOne(context.Background(), bson.M{"_id": "user-id"}).Decode(user)
	s.NoError(err)
	s.Equal("username", user.Username)
	s.Equal("user-id", user.ID)
}
func (s *userCacheSuite) TestUserCacheRepository_Find() {
	user := models.User{
		ID:       "user-id",
		Username: "username",
	}

	_, err := s.getCollection().InsertOne(context.Background(), user, &options.InsertOneOptions{})
	s.NoError(err)

	newUser, err := s.repo.Find(context.Background(), user.ID)
	s.NoError(err)
	s.NotNil(newUser)
	s.Equal(user.ID, newUser.ID)
}

func (s *userCacheSuite) TestUserCacheRepository_FindFromFallback() {
	s.mock.On("Find", mock.Anything, "user-id").Return(&models.User{
		ID:       "user-id",
		Username: "username",
	}, nil)

	newUser, err := s.repo.Find(context.Background(), "user-id")
	s.NoError(err)
	s.NotNil(newUser)
	s.Equal("user-id", newUser.ID)
	s.Equal("username", newUser.Username)
}

func (s *userCacheSuite) getCollection() *mongo.Collection {
	return s.db.Database(testDBName).Collection(testCollectionName)
}
