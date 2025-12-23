package cartservice

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"

func (s Service) UpdateItemQuantity(cartItemID int, act string) error {
	const op = "usecase.UpdateItemQuantity"

	if act != "add" && act != "sub" {
		return richerror.New(op).WithMessage("invalid action")
	}

	if err := s.repo.UpdateQuantity(cartItemID, act); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
