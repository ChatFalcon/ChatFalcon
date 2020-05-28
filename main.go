package main

import (
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/installkey"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/labstack/echo"
	"github.com/plutov/echo-logrus"
	"github.com/sirupsen/logrus"
	"net/http"
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

	// TODO: Add SSR, make this multiple routes, guess the status when serving.
	router.Router.GET("/", func(c echo.Context) error {
		t, err := c.Get("config").(*config.ServerConfig).RenderThemeHTML("https://"+c.Request().Host+"/", "Home")
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, t)
	})

	// Initialise the client.
	Host := os.Getenv("HOST")
	if Host == "" {
		Host = "127.0.0.1:8080"
	}
	logrus.Fatal(router.Router.Start(Host))
}
