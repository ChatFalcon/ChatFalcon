package main

import (
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"time"
)

func reqLog(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		StartTime := time.Now()
		defer func() {
			r := context.Request()
			logrus.Info(context.Response().Status, " ", r.Method, " ", r.RequestURI, " ", time.Now().Sub(StartTime))
		}()
		return handlerFunc(context)
	}
}

func init() {
	router.Router.Use(reqLog)
}
