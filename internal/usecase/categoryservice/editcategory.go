package categoryservice

import (
	"context"
	"strings"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) EditCategory(
	ctx context.Context,
	id int,
	req param.CategoryAddRequest,
) error {

	const op = "categoryservice.EditCategory"

	if id <= 0 {
		return richerror.New(op).
			WithMessage("invalid category id")
	}

	if strings.TrimSpace(req.Name) == "" {
		return richerror.New(op).
			WithMessage("category name is required")
	}

	if strings.TrimSpace(req.Slug) == "" {
		return richerror.New(op).
			WithMessage("category slug is required")
	}

	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	cat := entity.Category{
		ID:   id,
		Name: req.Name,
		Slug: req.Slug,
	}

	if err := s.repo.UpdateCategory(ctx, cat); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
