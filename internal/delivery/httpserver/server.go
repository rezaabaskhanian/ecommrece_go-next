package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/carthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/categoryhandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/checkouthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/producthandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/delivery/httpserver/userhandler"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/cartservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/categoryservice"
	checkoutservcie "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/checkoutservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/productservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/userservice"
)

type Server struct {
	port     int
	user     userhandler.Handler
	product  producthandler.Handler
	cart     carthandler.Handler
	checkout checkouthandler.Handler
	category categoryhandler.Handler
}

func New(port int, authSvc authservice.Service, userSvc userservice.Service,
	authConfig authservice.Config, productSvc productservice.Service,
	cartSvc cartservice.Service,
	checkSvc checkoutservcie.Service, categorySvc categoryservice.Service) Server {
	return Server{
		port:    port,
		user:    userhandler.New(userSvc, authSvc, authConfig),
		product: producthandler.New(productSvc, authConfig, authSvc),

		cart: carthandler.New(cartSvc, authSvc, authConfig),

		checkout: checkouthandler.New(checkSvc),

		category: categoryhandler.New(categorySvc, authSvc, authConfig),
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

	s.category.SetCategoryRoutes(e)

	// Start server
	// TODO : handle config for httpserver
	// e.Logger.Fatal(e.Start(":8084"))

	addr := fmt.Sprintf(":%d", s.port)
	e.Logger.Infof("listening on %s", addr)

	e.Logger.Fatal(e.Start(addr))
}
