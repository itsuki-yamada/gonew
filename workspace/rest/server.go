package rest

import (
	"context"
	"log/slog"

	log "github.com/itsuki-yamada/gonew/workspace/log"
	rest "github.com/itsuki-yamada/gonew/workspace/rest/handler/systems/ping"
	"github.com/itsuki-yamada/gonew/workspace/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				log.Logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	pingHandler := rest.NewPingHandler(&usecase.FetchPingUsecase{})
	g := e.Group("/systems")
	g.GET("/ping", pingHandler.Get)

	e.Logger.Fatal(e.Start(":8080"))
}
