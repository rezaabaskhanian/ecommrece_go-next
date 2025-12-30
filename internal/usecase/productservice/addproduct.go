package productservice

import (
	"context"
	"strings"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) AddProduct(
	ctx context.Context,
	p entity.Product,
) error {

	const op = "productservice.AddProduct"

	if strings.TrimSpace(p.Name) == "" {
		return richerror.New(op).WithMessage("product name is required")
	}

	if p.Price <= 0 {
		return richerror.New(op).WithMessage("invalid price")
	}

	if p.CategoryID <= 0 {
		return richerror.New(op).WithMessage("invalid category")
	}

	if err := s.repo.AddProduct(ctx, p); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
