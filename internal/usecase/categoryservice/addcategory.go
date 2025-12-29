package categoryservice

import (
	"context"
	"strings"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) AddCategory(ctx context.Context, req param.CategoryAddRequest) error {
	const op = "categoryservice.AddCategory"

	if strings.TrimSpace(req.Name) == "" {
		return richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage("category name is required")
	}

	if strings.TrimSpace(req.Slug) == "" {
		return richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage("category slug is required")
	}

	var Id int

	cat := entity.Category{
		ID:   Id,
		Name: req.Name,
		Slug: req.Slug,
	}

	err := s.repo.AddCategory(ctx, cat)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
