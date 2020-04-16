package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	General *GeneralHandler
	Home    *HomeHandler
	Payme   *PaymeHandler
	Blog    *BlogHandler
}

func NewHandler(e *echo.Echo) *Handler {
	return &Handler{
		General: NewGeneralHandler(),
		Home:    NewHomeHandler(),
		Payme:   NewPaymeHandler(),
		Blog:    NewBlogHandler(),
	}
}
