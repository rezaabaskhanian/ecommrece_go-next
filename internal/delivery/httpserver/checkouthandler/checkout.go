package checkouthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

func getClaims(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) CheckOutCart(c echo.Context) error {
	const op = "checkouthandler.CheckOutCart"

	claims := getClaims(c)
	userID := int(claims.UserID)

	result, err := h.checkoutSvc.Checkout(userID)
	if err != nil {
		// اینجا می‌تونی error mapping انجام بدی
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
			"op":    op,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "checkout successful",
		"order":   result.Order,
		"items":   result.Items,
	})
}
