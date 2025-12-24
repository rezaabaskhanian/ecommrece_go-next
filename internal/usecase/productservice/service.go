package productservice

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
)

type Repository interface {
	ShowAll(page, limit int) ([]entity.Product, int, error)
	DecreaseStock(ctx context.Context, tx pgx.Tx, productID, qty int) error
	GetProductWithID(ID int) (entity.Product, error)
	Search(q string, page int) ([]entity.Product, int, error)

	// AddProduct(productID int, name, desc string, price float64, stock int, category string, imageurl string) error
	// EditProduct(productID int, name, desc string, price float64, stock int, category string, imageurl string) error
	// DeleteProduct(productID int) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

// GetProductWithID implements checkoutservcie.ProductRepository.
func (s Service) GetProductWithID(ID int) (entity.Product, error) {
	return s.repo.GetProductWithID(ID)
}

// DecreaseStock implements checkoutservcie.ProductRepository.
func (s Service) DecreaseStock(ctx context.Context, tx pgx.Tx, productID int, qty int) error {
	return s.repo.DecreaseStock(ctx, tx, productID, qty)
}
