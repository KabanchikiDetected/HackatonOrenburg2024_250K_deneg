package mongo

import (
	"context"
	"fmt"

	"github.com/KabanchikiDetected/users/internal/config"
	"github.com/KabanchikiDetected/users/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "users"
)

type Storage struct {
	col *mongo.Collection
}

func New() *Storage {
	cfg := config.Config().Storage
	client := connect()
	col := client.Database(cfg.DBName).Collection(collectionName)
	return &Storage{
		col: col,
	}
}

func (s *Storage) SaveUser(ctx context.Context, user models.User) error {
	_, err := s.col.InsertOne(ctx, user)
	if err != nil {
		fmt.Print(err.Error())
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {

	var user models.User
	err := s.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func connect() *mongo.Client {
	opts := createOptions()
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	return client

}

func createOptions() *options.ClientOptions {
	cfg := config.Config().Storage
	cred := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	opts := options.Client()
	opts.ApplyURI(cfg.URL)
	opts.SetAuth(cred)
	return opts
}
