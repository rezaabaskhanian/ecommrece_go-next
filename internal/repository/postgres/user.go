package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

type Repository struct {
	DB *pgxpool.Pool
}

func MyNewPost(db *pgxpool.Pool) UserRepository {
	return &Repository{DB: db}
}

// GetUserByID implements UserRepository.
func (r *Repository) GetUserByID(userID int) (entity.User, error) {
	panic("unimplemented")
}

// GetUserByPhoneNumber implements UserRepository.
func (r *Repository) GetUserByPhoneNumber(phoneNmber string) (entity.User, error) {
	panic("unimplemented")
}

// Register implements UserRepository.
func (r *Repository) Register(u entity.User) (entity.User, error) {
	const op = "postgres.Register"

	query := `INSERT INTO Urls (name,phone_number,password ,created_at) VALUES ($1,$2,$3,NOW()) RETURNING id`

	var id uint
	err := r.DB.QueryRow(context.Background(), query, u.Name, u.PhoneNumber, u.Password).Scan(&id)

	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound)
	}
	u.ID = id
	return u, nil
}
