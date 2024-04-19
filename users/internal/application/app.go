package application

import (
	"log/slog"

	"github.com/KabanchikiDetected/users/internal/server"
	"github.com/KabanchikiDetected/users/internal/service"
	"github.com/KabanchikiDetected/users/internal/storage/mongo"
)

type App struct {
	server *server.Server
}

func New(log *slog.Logger) *App {
	storage := mongo.New()
	service := service.New(log, storage)
	server := server.New(service)
	return &App{
		server: server,
	}
}

func (a *App) Run() {
	go a.server.Run()
}

func (a *App) Stop() {
	a.server.Stop()
}
