package producthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h HandlerProdcut) ShowAllProduct(c echo.Context) error {
	const op = "producthandler.ShowAllProduct"

	res, err := h.productSvc.ShowAllProduct()

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, res)

}
