package categoryhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
)

func (h Handler) AddCategory(c echo.Context) error {

	const op = "categoryhandler.AddCategory"

	var req param.CategoryAddRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request body",
		})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid Name",
		})
	} else if req.Slug == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid Slug",
		})
	}

	err := h.categorysvc.AddCategory(
		c.Request().Context(),
		req,
	)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// 4️⃣ پاسخ موفق
	return c.JSON(http.StatusOK, echo.Map{
		"message": "category updated successfully",
	})

}
