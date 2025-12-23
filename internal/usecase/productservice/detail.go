package productservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) ProductDetail(ID int) (param.Product, error) {

	const op = "productservice.ProductDetail"

	res, err := s.repo.GetProductWithID(ID)

	if err != nil {
		return param.Product{}, richerror.New(op).WithErr(err)
	}

	return param.Product{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Stock:       res.Stock,
		Category:    res.Category,
		ImageURL:    res.ImageURL,
	}, nil
}
