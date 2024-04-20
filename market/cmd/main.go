package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/application"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/config"
)

func main() {
	config.LoadConfig()

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(h)

	app := application.New(log)
	app.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	app.Stop()
}
