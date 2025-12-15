package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/producthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/userhandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/productservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/userservice"
)

type Server struct {
	userhandler userhandler.Handler

	producthandler producthandler.HandlerProdcut
}

func New(authSvc authservice.Service, userSvc userservice.Service, authConfig authservice.Config, productSvc productservice.Service) Server {
	return Server{
		userhandler:    userhandler.New(userSvc, authSvc, authConfig),
		producthandler: producthandler.New(productSvc, authConfig, authSvc),
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.userhandler.SetUserRoutes(e)

	s.producthandler.SetProductRoutes(e)

	// Start server
	// TODO : handle config for httpserver
	e.Logger.Fatal(e.Start(":8084"))
}
