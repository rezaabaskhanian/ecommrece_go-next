package producthandler

import "github.com/labstack/echo/v4"

func (h Handler) SetProductRoutes(e *echo.Echo) {

	productGroup := e.Group("product")

	productGroup.GET("/showAllProduct", h.ShowAllProduct)

	productGroup.GET("/detail/:id", h.GetDetailProduct)

	productGroup.GET("/search", h.SearchProduct)

}
