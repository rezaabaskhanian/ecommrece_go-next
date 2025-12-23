package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/carthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/checkouthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/producthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/userhandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/cartservice"
	checkoutservcie "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/checkoutservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/productservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/userservice"
)

type Server struct {
	user     userhandler.Handler
	product  producthandler.Handler
	cart     carthandler.Handler
	checkout checkouthandler.Handler
}

func New(authSvc authservice.Service, userSvc userservice.Service,
	authConfig authservice.Config, productSvc productservice.Service,
	cartSvc cartservice.Service,
	checkSvc checkoutservcie.Service) Server {
	return Server{
		user:    userhandler.New(userSvc, authSvc, authConfig),
		product: producthandler.New(productSvc, authConfig, authSvc),

		cart: carthandler.New(cartSvc, authSvc, authConfig),

		checkout: checkouthandler.New(checkSvc),
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.user.SetUserRoutes(e)

	s.product.SetProductRoutes(e)

	s.cart.SetCartRoutes(e)

	s.checkout.SetCheckRoutes(e)

	// Start server
	// TODO : handle config for httpserver
	e.Logger.Fatal(e.Start(":8084"))
}
