package producthandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h HandlerProdcut) GetDetailProduct(c echo.Context) error {

	const op = "producthandler.GetDetailProduct"

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("invalid product id")
	}

	res, err := h.productSvc.ProductDetail(productID)

	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("")
	}

	return c.JSON(http.StatusOK, res)
}
