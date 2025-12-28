package productservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) Search(q string, page int) (param.PaginateResponse, error) {
	const op = "productservice.search"

	products, totalItems, err := s.repo.Search(q, page)

	if err != nil {
		return param.PaginateResponse{}, richerror.New(op).WithErr(err)
	}

	res := make([]param.Product, 0, len(products))

	for _, p := range products {
		res = append(res, param.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			CategoryID:  p.CategoryID,
			ImageURL:    p.ImageURL,
		})
	}

	limit := 10

	totalPage := (totalItems + limit - 1) / limit

	return param.PaginateResponse{
		Products:    res,
		CurrentPage: page,
		TotalItems:  totalItems,
		TotalPage:   totalPage,
	}, nil

}
