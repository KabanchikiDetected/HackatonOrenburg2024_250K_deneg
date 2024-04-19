package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/KabanchikiDetected/users/internal/application"
	"github.com/KabanchikiDetected/users/internal/config"
)

func main() {
	config.LoadConfig()

	log := slog.Default()

	app := application.New(log)
	app.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	app.Stop()

}
