package productservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) ShowAllProduct(req param.PaginateRequest) (param.PaginateResponse, error) {

	const op = "productservice.ShowAllProduct"

	limit := 10 // تعداد محصول در هر صفحه
	products, totalItems, err := s.repo.ShowAll(req.Page, limit)

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
			Category:    p.Category,
			ImageURL:    p.ImageURL,
			CreatedAt:   p.CreatedAt,
		})
	}
	totalPage := (totalItems + limit - 1) / limit

	return param.PaginateResponse{
		Products:    res,
		CurrentPage: req.Page,
		TotalPage:   totalPage,
		TotalItems:  totalItems,
	}, nil

}
