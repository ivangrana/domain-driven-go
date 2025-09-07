
package persistence

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"domain-driven-go/src/internal/domain/model"
	"domain-driven-go/src/internal/domain/repository"
)

// MongoDBUserRepository is a MongoDB implementation of the UserRepository.
type MongoDBUserRepository struct {
	collection *mongo.Collection
}

// NewMongoDBUserRepository creates a new MongoDBUserRepository.
func NewMongoDBUserRepository(connectionString, dbName, collectionName string) (repository.UserRepository, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBUserRepository{collection: collection}, nil
}

// Save saves a user to the repository.
func (r *MongoDBUserRepository) Save(user *model.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

// FindByID finds a user by their ID.
func (r *MongoDBUserRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
