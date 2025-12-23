package cartservice

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"

func (s Service) RemoveItem(cartItemID int) error {
	const op = "usecase.RemoveItem"

	if err := s.repo.RemoveItem(cartItemID); err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}
