package application

import (
	"context"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/students/internal/server"
	"github.com/KabanchikiDetected/hackaton/students/internal/service"
	mongoStorage "github.com/KabanchikiDetected/hackaton/students/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	log        *slog.Logger
	server     *server.Server
	collection *mongo.Collection
}

func New(log *slog.Logger) *App {
	collection := connectToMongoDB()
	usersStorage := mongoStorage.New(collection)

	usersService := service.New(log, usersStorage)

	server := server.New(usersService)
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
