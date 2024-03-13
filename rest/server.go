package rest

import (
	rest "github.com/itsuki-yamada/gonew/rest/handler/systems/ping"
	usecase "github.com/itsuki-yamada/gonew/usecase/systems/ping"
	"github.com/labstack/echo/v4"
)

func initRouting(e *echo.Echo) {
	pingHandler := rest.NewPingHandler(&usecase.PingUsecase{})
	g := e.Group("/systems")
	g.GET("/ping", pingHandler.Get)
}

func Run() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
