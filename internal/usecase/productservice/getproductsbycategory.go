package productservice

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) GetProductsByCategoryPaginated(ctx context.Context, req param.ProductWithCategoryReq) (param.PaginateResponse, error) {

	const op = "productservice.GetProductsByCategoryPaginated"

	limit := 10
	products, totalItems, err := s.repo.ShowByCategory(req.Slug, req.Page, limit)

	if err != nil {
		return param.PaginateResponse{}, richerror.New(op).WithErr(err)
	}

	res := make([]param.Product, 0, len(products))

	for _, p := range products {
		res = append(res, param.Product{

			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			Price:        p.Price,
			Stock:        p.Stock,
			ImageURL:     p.ImageURL,
			CategoryID:   p.CategoryID,
			CategoryName: p.CategoryName,
			CategorySlug: p.CategorySlug,
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
