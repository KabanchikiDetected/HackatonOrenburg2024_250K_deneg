package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/KabanchikiDetected/hackaton/students/internal/application"
)

func main() {
	log := slog.Default()
	app := application.New(log)

	app.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info("Shutting down...")

	app.Stop()
}
