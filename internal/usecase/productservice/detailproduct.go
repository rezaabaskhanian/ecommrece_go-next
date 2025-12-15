package productservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) ProductDetail(ID int) (param.ProductResponse, error) {

	const op = "productservice.ProductDetail"

	res, err := s.repo.GetProductWithID(ID)

	if err != nil {
		return param.ProductResponse{}, richerror.New(op).WithErr(err)
	}

	return param.ProductResponse{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Stock:       res.Stock,
		Category:    res.Category,
		ImageURL:    res.ImageURL,
		CreatedAt:   res.CreatedAt,
	}, nil
}
