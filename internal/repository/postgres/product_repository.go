package postgres

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type ProductRepository interface {
	ShowAll() ([]entity.Product, error)
	GetProductWithID(ID int) (entity.Product, error)
}
