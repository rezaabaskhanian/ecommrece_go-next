package categoryservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) GetList(limit int) ([]entity.Category, error) {

	const op = "categoryservice.GetList"

	res, err := s.repo.GetList(context.Background(), limit)

	if err != nil {
		return []entity.Category{}, richerror.New(op).WithErr(err)
	}

	return res, nil

}

func (s Service) GetCategoryWithName(name string) (entity.Category, error) {

	const op = "categoryservice.GetCategoryWithID"

	res, err := s.repo.GetByName(context.Background(), name)

	if err != nil {
		return entity.Category{}, richerror.New(op).WithErr(err)
	}

	return res, nil

}
