package carthandler

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/cartservice"
)

type Handler struct {
	cartSvc cartservice.Service

	authsvc    authservice.Service
	authConfig authservice.Config
}

func New(cartSvc cartservice.Service, authSvc authservice.Service, authConfig authservice.Config) Handler {
	return Handler{cartSvc: cartSvc,

		authsvc:    authSvc,
		authConfig: authConfig,
	}
}
