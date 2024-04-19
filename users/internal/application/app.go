package application

import (
	"log/slog"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/server"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/service"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/storage/mongo"
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
