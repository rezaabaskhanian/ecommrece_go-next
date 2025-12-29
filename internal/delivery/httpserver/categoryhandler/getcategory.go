package categoryhandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) GetCategryList(c echo.Context) error {

	const op = "categoryhandler.GetCategryList"
	limitStr := c.QueryParam("limit")

	limit := 4 // default
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid limit",
			})
		}
		limit = l
	}

	res, err := h.categorysvc.GetList(limit)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, res)

}
