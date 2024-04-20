package mongo

import (
	"context"
	"fmt"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/config"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user.ID = primitive.NewObjectID()
	// _, err := s.col.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user}, options.Update().SetUpsert(true))
	_, err := s.col.InsertOne(ctx, user)
	if err != nil {
		fmt.Print(err.Error())
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (s *Storage) MakeDeputy(ctx context.Context, id string) error {
	user, err := s.User(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	user.Role = models.Deputy
	_, err = s.col.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (s *Storage) User(ctx context.Context, id string) (models.User, error) {
	var user models.User

	oid, _ := primitive.ObjectIDFromHex(id)
	err := s.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func (s *Storage) UserByEmail(ctx context.Context, email string) (models.User, error) {
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
