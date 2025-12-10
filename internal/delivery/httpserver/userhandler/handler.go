package userhandler

import (
	usecasse "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase"
)

type Handler struct {
	usersvc usecasse.Service
}

func New(userSvc usecasse.Service) Handler {
	return Handler{
		usersvc: userSvc,
	}
}
