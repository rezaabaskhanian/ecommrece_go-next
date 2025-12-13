package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) userLogin(c echo.Context) error {

	const op = "userhandler.userLogin"

	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return richerror.New(op).WithErr(err)
	}

	res, err := h.usersvc.Login(req)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, res)

}
