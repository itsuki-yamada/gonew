package main

import (
	handler "github.com/itsuki-yamada/gonew/rest/handler/systems/ping"
	usecase "github.com/itsuki-yamada/gonew/usecase/systems/ping"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	h := handler.NewPingHandler(&usecase.PingUsecase{})
	g := e.Group("/systems")
	g.GET("/ping", h.Get)

	e.Logger.Fatal(e.Start(":8080"))
}
