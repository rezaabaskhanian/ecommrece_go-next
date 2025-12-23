package producthandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) ShowAllProduct(c echo.Context) error {
	const op = "producthandler.ShowAllProduct"

	pageStr := c.QueryParam("page")
	page := 1
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	req := param.PaginateRequest{Page: page}

	res, err := h.productSvc.ShowAllProduct(req)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, res)

}
