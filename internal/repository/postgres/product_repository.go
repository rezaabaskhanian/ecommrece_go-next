package postgres

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type ProductRepository interface {
	ShowAll(page, limit int) ([]entity.Product, int, error)
	GetProductWithID(ID int) (entity.Product, error)
	Serach(q string, page int) ([]entity.Product, int, error)
}
