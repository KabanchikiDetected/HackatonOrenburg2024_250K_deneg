package application

import (
	"context"
	"time"

	"github.com/KabanchikiDetected/hackaton/students/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB() *mongo.Collection {
	cfg := config.Config().Storage
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	opts := createOptions()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	db := client.Database(cfg.DBName)
	eventsCollection := db.Collection(cfg.Collection)
	return eventsCollection
}

func createOptions() *options.ClientOptions {
	cfg := config.Config().Storage
	cred := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	opts := options.Client()
	opts.ApplyURI(cfg.URI)
	opts.SetAuth(cred)
	return opts
}
