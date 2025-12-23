package checkouthandler

import "github.com/labstack/echo/v4"

func (h Handler) SetCheckRoutes(e *echo.Echo) {

	checkoutGroupe := e.Group("checkout")

	checkoutGroupe.GET("", h.CheckOutCart)
}
