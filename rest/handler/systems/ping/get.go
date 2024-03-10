package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *PingHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	return c.String(http.StatusOK, h.usecase.FetchPing(ctx))
}
