package producthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/middleware"
)

func (h Handler) SetProductRoutes(e *echo.Echo) {

	productGroup := e.Group("product")

	productGroup.GET("/showAllProduct", h.ShowAllProduct)

	productGroup.GET("/detail/:id", h.GetDetailProduct)

	productGroup.GET("/search", h.SearchProduct)

	// add this method

	productGroup.POST("/add", h.GetDetailProduct, middleware.Auth(h.authsvc, h.authConfig), middleware.RequireRole("admin"))
	productGroup.PUT("/update/{id}", h.GetDetailProduct, middleware.Auth(h.authsvc, h.authConfig), middleware.RequireRole("admin"))
	productGroup.DELETE("/delete/{id}", h.GetDetailProduct, middleware.Auth(h.authsvc, h.authConfig), middleware.RequireRole("admin"))

}
