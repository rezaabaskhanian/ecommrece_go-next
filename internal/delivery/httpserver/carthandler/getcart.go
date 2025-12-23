package carthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

func getClaims(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) GetCart(c echo.Context) error {
	const op = "carthandler.GetCart"

	claims := getClaims(c)

	res, err := h.cartSvc.GetCart(
		int(claims.UserID),
	)

	if err != nil {
		return richerror.New(op).
			WithKind(richerror.KindNotFound).
			WithMessage("سبد خرید یافت نشد")
	}

	return c.JSON(http.StatusOK, res)
}
