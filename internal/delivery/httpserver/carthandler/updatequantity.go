package carthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) UpdateQuantity(c echo.Context) error {
	const op = "carthandler.UpdateQuantity"

	var req param.CartUpdateQuantityRequest

	if err := c.Bind(&req); err != nil {
		return richerror.New(op).WithErr(err).WithMessage("Invalid request body")
	}

	err := h.cartSvc.UpdateItemQuantity(req.CartItemID, req.Action)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Item update from cart successfully",
	})
}
