package main

import (
	"github.com/bilfash/protoman/application"
	"github.com/bilfash/protoman/log"
)

func main() {
	app, err := application.New()
	if err != nil {
		logFatal(err, "unable to initiate protoman app")
	}
	start(app)
	defer stop(app)
}

func start(app application.App) {
	if err := app.Start(); err != nil {
		logFatal(err, "unable to start protoman app")
	}
}

func stop(app application.App) {
	if err := app.Stop(); err != nil {
		logFatal(err, "unable to stop protoman app")
	}
}

func logFatal(err error, message string) {
	log.Get().WithError(err).Fatal(message)
}
