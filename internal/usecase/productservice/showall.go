package productservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) ShowAllProduct() ([]param.ProductResponse, error) {

	const op = "productservice.ShowAllProduct"
	products, err := s.repo.ShowAll()

	if err != nil {
		return []param.ProductResponse{}, richerror.New(op).WithErr(err)
	}

	res := make([]param.ProductResponse, 0, len(products))

	for _, p := range products {
		res = append(res, param.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Category:    p.Category,
			ImageURL:    p.ImageURL,
			CreatedAt:   p.CreatedAt,
		})
	}

	return res, nil

}
