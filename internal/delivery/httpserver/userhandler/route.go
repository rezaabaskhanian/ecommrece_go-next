package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	userGroup := e.Group("users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	userGroup.POST("/resetpassword", h.UserResetPass)

	userGroup.GET("/profile", h.userProfile,
		middleware.Auth(h.authsvc, h.authConfig))

}
