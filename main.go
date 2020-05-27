package main

import (
	"github.com/ChatFalcon/ChatFalcon/installkey"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/plutov/echo-logrus"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/ChatFalcon/ChatFalcon/setup"
)

// TODO: Add favicon route.

func main() {
	// Handles debug logging.
	if os.Getenv("DEBUG") == "1" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Configure logrus.
	echologrus.Logger = logrus.New()
	router.Router.Logger = echologrus.GetEchoLogger()

	// Initialise the MongoDB client.
	mongo.InitClient()

	// Get the install key.
	installkey.GetInstallKey()

	// Initialise the client.
	Host := os.Getenv("HOST")
	if Host == "" {
		Host = "127.0.0.1:8080"
	}
	logrus.Fatal(router.Router.Start(Host))
}
