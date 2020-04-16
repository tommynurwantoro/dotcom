package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GeneralHandler struct{}

func (g *GeneralHandler) HealthzHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func NewGeneralHandler() *GeneralHandler {
	return &GeneralHandler{}
}
