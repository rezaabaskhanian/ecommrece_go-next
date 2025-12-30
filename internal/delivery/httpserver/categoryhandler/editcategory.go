package categoryhandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

func getClaims(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}
func (h Handler) EditCategory(c echo.Context) error {
	const op = "categoryhandler.EditCategory"

	var req param.CategoryAddRequest
	var id int

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid category id",
		})
	}
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

	// 3️⃣ صدا زدن service
	err = h.categorysvc.EditCategory(
		c.Request().Context(),
		id,
		req,
	)

	if err != nil {
		// اینجا می‌تونی error mapper داشته باشی
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// 4️⃣ پاسخ موفق
	return c.JSON(http.StatusOK, echo.Map{
		"message": "category updated successfully",
	})

}
