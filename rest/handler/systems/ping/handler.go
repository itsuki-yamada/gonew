package rest

import (
	"context"
	"net/http"

	"github.com/itsuki-yamada/gonew/domain"
	"github.com/labstack/echo/v4"
)

type (
	PingHandler struct {
		usecase PingUsecase
	}

	PingUsecase interface {
		FetchPing(c context.Context, ping domain.Pinger) domain.Message
	}
)

func NewPingHandler(u PingUsecase) *PingHandler {
	return &PingHandler{
		usecase: u,
	}
}

func (h *PingHandler) Get(c echo.Context) error {
	msg := c.QueryParam("msg")
	ping := domain.NewPinger(msg)

	ctx := c.Request().Context()
	resText := h.usecase.FetchPing(ctx, ping)
	return c.String(http.StatusOK, string(resText))
}
