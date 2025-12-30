package productservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) EditProduct(
	ctx context.Context,
	productID int,
	p entity.Product,
) error {

	const op = "productservice.EditProduct"

	if productID <= 0 {
		return richerror.New(op).WithMessage("invalid product id")
	}

	_, err := s.repo.GetProductWithID(ctx, productID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	p.ID = productID

	if err := s.repo.UpdateProduct(ctx, p); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
