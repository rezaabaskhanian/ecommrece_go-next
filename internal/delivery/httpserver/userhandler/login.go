package userhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) userLogin(c echo.Context) error {

	const op = "userhandler.userLogin"

	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		log.Println(op, "Bind error:", err)
		return richerror.New(op).WithErr(err)
	}

	if err := c.Validate(&req); err != nil {
		log.Println(op, "Validation error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.usersvc.Login(req)

	if err != nil {
		// تشخیص نوع خطا
		if err.Error() == "record Not Found" {
			log.Println(op, "User not found:", req.PhoneNumber)
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		log.Println(op, "Login service error:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	log.Println(op, "Login success:", req.PhoneNumber)

	return c.JSON(http.StatusOK, res)

}
