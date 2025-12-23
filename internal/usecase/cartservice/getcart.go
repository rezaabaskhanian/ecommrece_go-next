package cartservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) GetCart(userID int) (entity.CartWithItem, error) {
	const op = "usecase.GetCart"

	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return entity.CartWithItem{}, richerror.New(op).WithErr(err)
	}

	items, err := s.repo.GetCartItems(cart.ID)
	if err != nil {
		return entity.CartWithItem{}, richerror.New(op).WithErr(err)
	}

	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return entity.CartWithItem{
		Cart:       cart,
		Items:      items,
		TotalPrice: total,
		ItemCount:  len(items),
	}, nil
}
