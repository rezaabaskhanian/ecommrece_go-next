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

	const op = "chackouhandler.CheckOutCart"

	claims := getClaims(c)

	err := h.checkoutSvc.CheckOutOrder(int(claims.UserID))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"messge": "checkout is succeful",
	})

}
