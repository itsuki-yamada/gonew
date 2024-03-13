package main

import (
	"github.com/labstack/echo/v4"
)

var di *Container

func init() {
	di = NewDI()
}

func initRouting(e *echo.Echo) {
	g := e.Group("/systems")
	g.GET("/ping", di.PingHandler.Get)
}

func main() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
