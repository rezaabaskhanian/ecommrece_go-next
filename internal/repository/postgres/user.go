package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func MyNewPostgresUser(db *pgxpool.Pool) UserRepository {
	return &Repository{DB: db}
}

func (r *Repository) GetUserByID(userID int) (entity.User, error) {
	const op = "postgres.GetUserByID"

	query := `SELECT id, name, phone_number, password, avatar_url,created_at FROM users WHERE id = $1`

	var u entity.User

	err := r.DB.QueryRow(context.Background(), query, userID).Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &u.AvatarURL, &u.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		return entity.User{}, err
	}
	return u, nil

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

	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound)
	}
	u.ID = id
	return u, nil
}

func (r *Repository) ResetPassword(phoneNumber, hashedPassword string) error {
	const op = "postgres.ResetPassword"

	query := `UPDATE users SET password = $1 WHERE phone_number = $2`

	cmdTag, err := r.DB.Exec(context.Background(), query, hashedPassword, phoneNumber)

	fmt.Println("PhoneNumber:", phoneNumber)
	fmt.Println("HashedPassword:", hashedPassword)
	fmt.Println("RowsAffected:", cmdTag.RowsAffected())
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if cmdTag.RowsAffected() == 0 {
		return richerror.New(op).WithMessage("user not found")
	}

	return nil

}
