package application

import (
	"context"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/events/internal/server"
	eventsService "github.com/KabanchikiDetected/hackaton/events/internal/service/events"
	usersService "github.com/KabanchikiDetected/hackaton/events/internal/service/users"
	eventsMongoStorage "github.com/KabanchikiDetected/hackaton/events/internal/storage/mongo/events"
	usersMongoStorage "github.com/KabanchikiDetected/hackaton/events/internal/storage/mongo/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	log              *slog.Logger
	server           *server.Server
	eventsCollection *mongo.Collection
	usersCollection  *mongo.Collection
}

func New(log *slog.Logger) *App {
	eventsCollection, usersCollection := connectToMongoDB()
	eventsStorage := eventsMongoStorage.New(eventsCollection)
	usersStorage := usersMongoStorage.New(usersCollection, eventsCollection)

	eventService := eventsService.New(log, eventsStorage)
	usersService := usersService.New(log, usersStorage)

	server := server.New(eventService, usersService)
	return &App{
		log:              log,
		server:           server,
		eventsCollection: eventsCollection,
		usersCollection:  usersCollection,
	}
}

func (app *App) Start() {
	go app.server.Start()
}

func (app *App) Stop() {
	app.server.Stop()
	app.eventsCollection.Database().Client().Disconnect(context.Background())
	app.usersCollection.Database().Client().Disconnect(context.Background())
}
