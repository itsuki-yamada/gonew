package rest

import (
	rest "github.com/itsuki-yamada/gonew/rest/handler/systems/ping"
	"github.com/itsuki-yamada/gonew/usecase"
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	pingHandler := rest.NewPingHandler(&usecase.FetchPingUsecase{})
	g := e.Group("/systems")
	g.GET("/ping", pingHandler.Get)

	e.Logger.Fatal(e.Start(":8080"))
}
