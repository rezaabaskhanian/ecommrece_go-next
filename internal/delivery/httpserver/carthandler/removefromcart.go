package carthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) RemoveFromCart(c echo.Context) error {
	const op = "carthandler.RemoveFromCart"

	var req param.CartRemoveRequest

	if err := c.Bind(&req); err != nil {
		return richerror.New(op).WithErr(err).WithMessage("Invalid request body")
	}

	err := h.cartSvc.RemoveItem(req.CartItemID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Item remove from cart successfully",
	})

}
