package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

func getClaims(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) userProfile(c echo.Context) error {

	const op = "userhandler.userProfile"

	// var req param.ProfileRequest

	// if err := c.Bind(&req); err != nil {
	// 	return richerror.New(op).WithErr(err)
	// }

	claims := getClaims(c)

	res, err := h.usersvc.Profile(param.ProfileRequest{UserID: claims.UserID})

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return c.JSON(http.StatusOK, res)

}
