package producthandler

import "github.com/labstack/echo/v4"

func (h HandlerProdcut) SetProductRoutes(e *echo.Echo) {

	userGroup := e.Group("product")

	userGroup.GET("/showAllProduct", h.ShowAllProduct)

	userGroup.GET("/detail/:id", h.GetDetailProduct)

	userGroup.GET("/search", h.SearchProduct)

}
