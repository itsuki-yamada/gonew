package rest

import (
	"net/http"

	"github.com/itsuki-yamada/gonew/domain"
	"github.com/labstack/echo/v4"
)

func (h *PingHandler) Get(c echo.Context) error {
	msg := c.QueryParam("msg")
	ping := domain.NewPinger(msg)

	ctx := c.Request().Context()
	resText := h.usecase.Exec(ctx, ping)
	return c.String(http.StatusOK, string(resText))
}
