package mongo

import (
	"context"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/models"
	"github.com/stackus/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserCacheRepository struct {
	dbName    string
	tableName string
	db        *mongo.Client
	fallback  app.UserRepository
}

var _ app.UserCacheRepository = (*UserCacheRepository)(nil)

func NewUserCacheRepository(dbName, tableName string, db *mongo.Client, fallback app.UserRepository) UserCacheRepository {
	return UserCacheRepository{
		dbName:    dbName,
		tableName: tableName,
		db:        db,
		fallback:  fallback,
	}
}

func (r UserCacheRepository) Add(ctx context.Context, userID, username string) error {
	user := models.User{
		ID:       userID,
		Username: username,
	}

	_, err := r.getCollection().InsertOne(ctx, user, &options.InsertOneOptions{})

	return err
}

func (r UserCacheRepository) Find(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{
		ID: userID,
	}

	err := r.getCollection().FindOne(ctx, bson.M{"_id": userID}).Decode(user)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Wrap(err, "scanning user")
		}
		user, err = r.fallback.Find(ctx, userID)
		if err != nil {
			return nil, errors.Wrap(err, "user callback failed")
		}
		// attempt to add it to the cache
		return user, r.Add(ctx, user.ID, user.Username)
	}

	return user, nil
}

func (r UserCacheRepository) getCollection() *mongo.Collection {
	return r.db.Database(r.dbName).Collection(r.tableName)
}
