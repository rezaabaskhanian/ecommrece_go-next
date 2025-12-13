package userservice

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type Repository interface {
	GetUserByPhoneNumber(phoneNmber string) (entity.User, error)
	Register(u entity.User) (entity.User, error)
	GetUserByID(userID int) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(auth AuthGenerator, repo Repository) Service {
	return Service{auth: auth, repo: repo}
}
