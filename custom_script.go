package main

import (
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/labstack/echo"
)

func init() {
	router.Router.GET("/custom_script.js", func(c echo.Context) error {
		c.Response().Status = 200
		c.Response().Header().Set("Content-Type", "application/javascript")
		_, err := c.Response().Write([]byte(c.Get("config").(*config.ServerConfig).CustomScript))
		if err != nil {
			return err
		}
		return nil
	})
}
