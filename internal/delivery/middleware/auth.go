package middleware

import (
	"net/http"

	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	cfg "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	const op = "middleware.Auth"

	// get token without "Bearer "  echo is delete berare then
	return mw.WithConfig(mw.Config{
		ContextKey: cfg.AuthMiddlewareContextKey,
		SigningKey: []byte(config.SignKey),
		// TODO - as sign method string to config...
		SigningMethod: "HS256",

		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "لطفا ابتدا وارد حساب کاربری خود شوید",
			})
		},

		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {

			claims, err := service.ParseToken("Bearer " + auth)

			if err != nil {
				return nil, richerror.New(op).WithErr(err).WithMessage("dont create clamis")
			}

			return claims, nil
		},
	})
}

func RequireRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			claims, ok := c.Get(cfg.AuthMiddlewareContextKey).(*authservice.Claims)
			if !ok || claims == nil {
				return echo.ErrUnauthorized
			}

			for _, role := range roles {
				if claims.Role == role {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "شما اجازه دسترسی به این بخش را ندارید",
			})
		}
	}
}
