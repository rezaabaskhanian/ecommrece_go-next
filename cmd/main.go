package main

import (
	"time"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/repository/postgres"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/productservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/userservice"
)

const (
	jwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24 * 7
	RefreshTokenExpireDuration = time.Hour * 24 * 7 * 30
)

const ()

func main() {

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8081},
		Auth: authservice.Config{
			SignKey:              jwtSignKey,
			AccessExpirationTime: AccessTokenExpireDuration,

			RefreshExpiratoonTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		MyPostgres: postgres.Config{
			UserName: "user",
			Password: "pass",
			Host:     "localhost",
			Port:     5433,
			DBName:   "ecommerce",
		},
	}

	authSvc, userSvc, authConfig, productSvc := setupService(cfg)

	server := httpserver.New(authSvc, userSvc, authConfig, productSvc)

	server.Serve()

}

func setupService(cfg config.Config) (authservice.Service, userservice.Service, authservice.Config, productservice.Service) {

	authSvc := authservice.New(cfg.Auth)

	myPostgresRepo := postgres.New(cfg.MyPostgres)

	myPostgresRepoUser := postgres.MyNewPostgresUser(myPostgresRepo)

	myPostgresRepoProduct := postgres.MyNewPostgresProduct(myPostgresRepo)

	userSvc := userservice.New(authSvc, myPostgresRepoUser)

	productSvc := productservice.New(myPostgresRepoProduct)

	return authSvc, userSvc, cfg.Auth, productSvc

}
