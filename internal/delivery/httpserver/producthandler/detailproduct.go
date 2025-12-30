package producthandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
)

func (h Handler) GetDetailProduct(c echo.Context) error {

	const op = "producthandler.GetDetailProduct"

	pageStr := c.QueryParam("page")
	slug := c.QueryParam("slug")
	page := 1
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if slug == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "slug is required",
		})
	}

	req := param.ProductWithCategoryReq{Slug: slug, Page: page}

	res, err := h.productSvc.GetProductsByCategoryPaginated(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, res)
}
