package middleware

import (
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
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {

			claims, err := service.ParseToken("Bearer " + auth)

			if err != nil {
				return nil, richerror.New(op).WithErr(err).WithMessage("dont create clamis")
			}

			return claims, nil
		},
	})
}
