package carthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/middleware"
)

func (h Handler) SetCartRoutes(e *echo.Echo) {

	cartGroup := e.Group("cart")

	cartGroup.GET("/getcart", h.GetCart, middleware.Auth(h.authsvc, h.authConfig))
	cartGroup.GET("/add", h.AddToCart, middleware.Auth(h.authsvc, h.authConfig))
	cartGroup.GET("/remove", h.RemoveFromCart, middleware.Auth(h.authsvc, h.authConfig))

}
