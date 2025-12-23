package carthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) AddToCart(c echo.Context) error {
	const op = "carthandler.AddToCart"

	claims := getClaims(c)

	var req param.CartItemRequest
	if err := c.Bind(&req); err != nil {
		return richerror.New(op).WithErr(err).WithMessage("Invalid request body")
	}

	if req.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Quantity must be greater than zero"})
	}

	err := h.cartSvc.AddItem(int(claims.UserID), req.ProductID, req.Quantity)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Item added to cart successfully",
	})
}
