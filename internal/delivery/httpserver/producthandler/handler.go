package producthandler

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/productservice"
)

type HandlerProdcut struct {
	productSvc productservice.Service

	authsvc authservice.Service

	authConfig authservice.Config
}

func New(productSvc productservice.Service, authConfig authservice.Config, authsvc authservice.Service) HandlerProdcut {
	return HandlerProdcut{productSvc: productSvc, authConfig: authConfig, authsvc: authsvc}
}
