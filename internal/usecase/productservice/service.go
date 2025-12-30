package productservice

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/repository/model"
)

type Repository interface {
	ShowAll(page, limit int) ([]model.ProductWithCategory, int, error)
	DecreaseStock(ctx context.Context, tx pgx.Tx, productID, qty int) error
	GetProductWithID(ctx context.Context, ID int) (entity.Product, error)
	Search(q string, page int) ([]model.ProductWithCategory, int, error)

	ShowByCategory(categorySlug string, page int, limit int) ([]model.ProductWithCategory, int, error)

	AddProduct(ctx context.Context, p entity.Product) error
	UpdateProduct(ctx context.Context, p entity.Product) error
	DeleteProduct(ctx context.Context, productID int) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

// GetProductWithID implements checkoutservcie.ProductRepository.
func (s Service) GetProductWithID(ctx context.Context, ID int) (entity.Product, error) {
	return s.repo.GetProductWithID(ctx, ID)
}

// DecreaseStock implements checkoutservcie.ProductRepository.
func (s Service) DecreaseStock(ctx context.Context, tx pgx.Tx, productID int, qty int) error {
	return s.repo.DecreaseStock(ctx, tx, productID, qty)
}
