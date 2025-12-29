package main

import (
	"time"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/config"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/repository/postgres"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/cartservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/categoryservice"
	checkoutservcie "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/checkoutservice"
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

	authSvc, userSvc, authConfig, productSvc, cartSvc, checkoutSvc, categorySvc := setupService(cfg)

	server := httpserver.New(authSvc, userSvc, authConfig, productSvc, cartSvc, checkoutSvc, categorySvc)

	server.Serve()

}

func setupService(cfg config.Config) (authservice.Service, userservice.Service, authservice.Config,
	productservice.Service, cartservice.Service, checkoutservcie.Service, categoryservice.Service) {

	authSvc := authservice.New(cfg.Auth)

	myPostgresRepo := postgres.New(cfg.MyPostgres)

	userRepo := postgres.NewUserRepository(myPostgresRepo)
	productRepo := postgres.NewProductRepository(myPostgresRepo)
	cartRepo := postgres.NewCartRepository(myPostgresRepo)
	orderRepo := postgres.NewOrderRepository(myPostgresRepo)
	categoryRepo := postgres.NewCategoryRepository(myPostgresRepo)

	userSvc := userservice.New(authSvc, userRepo)
	productSvc := productservice.New(productRepo)
	cartSvc := cartservice.New(cartRepo)
	categorySvc := categoryservice.New(categoryRepo)

	checkoutSvc := checkoutservcie.New(cartRepo, productSvc, orderRepo, orderRepo)

	return authSvc, userSvc, cfg.Auth, productSvc, cartSvc, checkoutSvc, categorySvc

}
