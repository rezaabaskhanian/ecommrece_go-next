package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

type Repository struct {
	DB *pgxpool.Pool
}

func MyNewPostgresUser(db *pgxpool.Pool) UserRepository {
	return &Repository{DB: db}
}

// GetUserByID implements UserRepository.
func (r *Repository) GetUserByID(userID int) (entity.User, error) {
	return entity.User{}, errors.New("not implemented")

}

// GetUserByPhoneNumber implements UserRepository.
func (r *Repository) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "postgres.Register"
	query := `SELECT id, name, phone_number, password, created_at FROM users WHERE phone_number = $1`

	var u entity.User
	err := r.DB.QueryRow(context.Background(), query, phoneNumber).Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &u.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return entity.User{}, err
	}
	return u, nil

}

// Register implements UserRepository.
func (r *Repository) Register(u entity.User) (entity.User, error) {
	const op = "postgres.Register"

	query := `INSERT INTO users (name,phone_number,password ,created_at) VALUES ($1,$2,$3,NOW()) RETURNING id`

	var id uint
	err := r.DB.QueryRow(context.Background(), query, u.Name, u.PhoneNumber, u.Password).Scan(&id)

	fmt.Println(err, "postgres.Register")

	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound)
	}
	u.ID = id
	return u, nil
}
