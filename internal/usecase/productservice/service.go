package productservice

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type Repository interface {
	ShowAll(page, limit int) ([]entity.Product, int, error)
	GetProductWithID(ID int) (entity.Product, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
