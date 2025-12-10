package usecasse

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type Repository interface {
	GetUserByPhoneNumber(phoneNmber string) (entity.User, error)
	Register(u entity.User) (entity.User, error)
	GetUserByID(userID int) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
