package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

type (
	PingHandler struct {
		usecase PingUsecase
	}

	PingUsecase interface {
		FetchPing(c context.Context) string
	}
)

func NewPingHandler(u PingUsecase) *PingHandler {
	return &PingHandler{
		usecase: u,
	}
}

func (h *PingHandler) Routing(e *echo.Echo) {
	g := e.Group("/systems")

	g.GET("/ping", h.Get)
}
