package categoryservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
)

type Repository interface {
	GetList(ctx context.Context, limit int) ([]entity.Category, error)
	GetByName(ctx context.Context, name string) (entity.Category, error)
	GetByID(ctx context.Context, Id int) (entity.Category, error)

	AddCategory(ctx context.Context, c entity.Category) error
	UpdateCategory(ctx context.Context, c entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}

}
