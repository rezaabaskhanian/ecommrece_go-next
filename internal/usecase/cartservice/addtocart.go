package cartservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

// TODO : change addtocart to update cart

func (s Service) AddItem(userID, productID, quantity int) error {
	const op = "usecase.AddItem"

	if quantity <= 0 {
		return richerror.New(op).WithMessage("quantity must be greater than zero")
	}

	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if err := s.repo.AddItemToCart(cart.ID, productID, quantity); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
