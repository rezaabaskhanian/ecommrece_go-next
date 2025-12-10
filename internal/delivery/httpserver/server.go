package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/userhandler"
)

type Server struct {
	userhandler userhandler.Handler
}

func New(userhandler userhandler.Handler) Server {
	return Server{
		userhandler: userhandler,
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.userHandler.SetUserRoutes(e)

	// Start server
	// TODO : handle config for httpserver
	e.Logger.Fatal(e.Start(":8084"))
}
