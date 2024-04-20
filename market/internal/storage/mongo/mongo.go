package mongo

import (
	"context"
	"fmt"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/config"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "product"
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

func (s *Storage) CreateProduct(ctx context.Context, product *models.Product) error {
	const op = "mongo.CreateProduct"

	_, err := s.col.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) Products(ctx context.Context) ([]models.Product, error) {
	const op = "mongo.GetProducts"

	var products []models.Product
	cursor, err := s.col.Find(ctx, primitive.D{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return products, nil
}

func (s *Storage) Product(ctx context.Context, id string) (models.Product, error) {
	const op = "mongo.Product"

	var product models.Product
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, fmt.Errorf("%s: %w", op, err)
	}
	err = s.col.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return product, fmt.Errorf("%s: %w", op, err)
	}
	return product, nil
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
