package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
)

func (h Handler) userRegister(c echo.Context) error {

	var req param.RegisterRequest

	//TODO : add validator for phonenuber and password and etc ...

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := h.usersvc.Register(req)

	if err != nil {
		return
	}

}
