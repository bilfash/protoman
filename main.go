package main

import (
	"github.com/bilfash/protoman/application"
	"github.com/bilfash/protoman/log"
)

func main() {
	app := application.New()
	if err := app.Start(); err != nil {
		log.Get().WithError(err).Fatal("unable to start protoman app")
	}
	defer stop(app)
}

func stop(app application.App) {
	if err := app.Stop(); err != nil {
		log.Get().WithError(err).Fatal("unable to stop protoman app")
	}
}
