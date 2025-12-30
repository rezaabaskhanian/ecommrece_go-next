package categoryhandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) DeleteCategory(c echo.Context) error {

	const op = "categoryhandler.DeleteCategory"

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid category id",
		})
	}

	err = h.categorysvc.DeleteCategory(
		c.Request().Context(),
		id,
	)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "category deleted successfully",
	})

}
