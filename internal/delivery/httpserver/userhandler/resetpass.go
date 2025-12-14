package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) UserResetPass(c echo.Context) error {
	const op = "userhandler.UserResetPass"

	var req param.PasswordRequest

	if err := c.Bind(&req); err != nil {

		return richerror.New(op).WithErr(err).WithMessage("dont req")

	}

	errSvc := h.usersvc.ResetPassword(req)

	if errSvc != nil {
		return richerror.New(op).WithErr(errSvc)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "password reset successfully",
	})

}
