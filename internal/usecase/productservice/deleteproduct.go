package productservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) DeleteProduct(
	ctx context.Context,
	productID int,
) error {

	const op = "productservice.DeleteProduct"

	if productID <= 0 {
		return richerror.New(op).WithMessage("invalid product id")
	}

	if err := s.repo.DeleteProduct(ctx, productID); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
