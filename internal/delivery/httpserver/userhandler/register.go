package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (h Handler) userRegister(c echo.Context) error {

	const op = "userhandler.userRegister"

	var req param.RegisterRequest

	//TODO : add validator for phonenuber and password and etc ...

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := h.usersvc.Register(req)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusCreated, res)

}
