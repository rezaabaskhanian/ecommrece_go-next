package postgres

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type UserRepository interface {
	GetUserByPhoneNumber(phoneNmber string) (entity.User, error)
	Register(u entity.User) (entity.User, error)
	GetUserByID(userID int) (entity.User, error)
}
