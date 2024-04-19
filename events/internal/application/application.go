package application

import (
	"context"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/events/internal/server"
	"github.com/KabanchikiDetected/hackaton/events/internal/service"
	mongoStorage "github.com/KabanchikiDetected/hackaton/events/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	log        *slog.Logger
	server     *server.Server
	collection *mongo.Collection
}

func New(log *slog.Logger) *App {
	collection := connectToMongoDB()
	storage := mongoStorage.New(collection)

	eventService := service.New(log, storage)

	server := server.New(eventService)
	return &App{
		log:        log,
		server:     server,
		collection: collection,
	}
}

func (app *App) Start() {
	go app.server.Start()
}

func (app *App) Stop() {
	app.server.Stop()
	app.collection.Database().Client().Disconnect(context.Background())
}
