package userhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
)

func (h Handler) userRegister(c echo.Context) error {

	const op = "userhandler.userRegister"

	var req param.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.usersvc.Register(req)

	if err != nil {
		// لاگ دقیق خطا
		log.Println(op, "Register service error:", err)

		// تشخیص نوع خطا
		if err.Error() == "duplicate phone_number" {
			return echo.NewHTTPError(http.StatusConflict, "Phone number already registered")
		}

		if err.Error() == "hashing failed" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, res)

}
