package userhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	userGroupe := e.Group("users")

	userGroupe.POST("/register", h.userRegister)
}
