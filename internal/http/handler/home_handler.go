package handler

import (
	"net/http"
	"tommynurwantoro/dotcom/internal/domain"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	Data domain.HomeData
}

func (h *HomeHandler) Index(c echo.Context) error {
	cc := c.(*domain.CustomContext)
	h.Data.Tagline = cc.Config.HomeTagline

	return c.Render(http.StatusOK, "home", h.Data)
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{
		Data: domain.HomeData{
			PageTitle: "Home",
		},
	}
}
