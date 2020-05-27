package main

import (
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	router.Router.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			IsSetup := func(Path string) bool {
				return Path == "/setup/" || Path == "/setup/ui.js" || Path == "/setup/ui.js.map"
			}

			cfg, err := config.GetConfig()
			if err == nil {
				if IsSetup(context.Request().RequestURI) {
					return context.Redirect(http.StatusTemporaryRedirect, "/")
				}
				context.Set("config", cfg)
				return handlerFunc(context)
			} else if err != mongo.ErrNoDocuments {
				return reqLog(func(context echo.Context) error {
					return context.String(http.StatusInternalServerError, err.Error())
				})(context)
			}

			if !IsSetup(context.Request().RequestURI) {
				return reqLog(func(context echo.Context) error {
					return context.Redirect(http.StatusTemporaryRedirect, "/setup/")
				})(context)
			}

			return handlerFunc(context)
		}
	})
}
