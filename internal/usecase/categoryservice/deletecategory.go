package categoryservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) DeleteCategory(
	ctx context.Context,
	id int,
) error {

	const op = "categoryservice.DeleteCategory"

	if id <= 0 {
		return richerror.New(op).WithMessage("invalid category id")

	}

	// چک وجود category
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// // آیا محصولی با این category وجود دارد؟
	// count, err := s.repo.CountProducts(ctx, id)
	// if err != nil {
	// 	return richerror.New(op).WithErr(err)
	// }

	// if count > 0 {
	// 	return richerror.New(op).
	// 		WithKind(richerror.KindConflict).
	// 		WithMessage("category has products and cannot be deleted")
	// }

	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
