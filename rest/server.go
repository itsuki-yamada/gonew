package main

import (
	"net/http"

	handler "github.com/itsuki-yamada/cleanarchtecture/rest/handler/systems/ping"
	usecase "github.com/itsuki-yamada/cleanarchtecture/usecase/systems/ping"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	h := handler.NewPingHandler(&usecase.PingUsecase{})
	h.Routing(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
