package userhandler

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/userservice"
)

type Handler struct {
	usersvc userservice.Service
	authsvc authservice.Service
}

func New(userSvc userservice.Service, authSvc authservice.Service) Handler {
	return Handler{
		usersvc: userSvc,
		authsvc: authSvc,
	}
}
