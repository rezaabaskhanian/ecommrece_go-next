package categoryhandler

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/categoryservice"
)

type Handler struct {
	categorysvc categoryservice.Service

	authsvc    authservice.Service
	authconfig authservice.Config
}

func New(categorySvc categoryservice.Service, authSvc authservice.Service, authConfig authservice.Config) Handler {
	return Handler{
		categorysvc: categorySvc,
		authsvc:     authSvc,
		authconfig:  authConfig,
	}
}
