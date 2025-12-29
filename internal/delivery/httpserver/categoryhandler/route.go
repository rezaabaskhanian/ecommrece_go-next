package categoryhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/middleware"
)

func (h Handler) SetCategoryRoutes(e *echo.Echo) {

	cartGroup := e.Group("categories")

	cartGroup.GET("", h.GetCategryList)
	cartGroup.GET("", h.GetCategryList)

	cartGroup.POST("/add", h.AddCategory, middleware.Auth(h.authsvc, h.authconfig), middleware.RequireRole("admin"))
	cartGroup.PUT("/edit/{id}", h.EditCategory, middleware.Auth(h.authsvc, h.authconfig), middleware.RequireRole("admin"))
	cartGroup.DELETE("/delete/{id}", h.DeleteCategory, middleware.Auth(h.authsvc, h.authconfig), middleware.RequireRole("admin"))

}
