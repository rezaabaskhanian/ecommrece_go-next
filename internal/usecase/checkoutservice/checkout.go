package checkoutservcie

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"

func (s Service) CheckOutOrder(userID int) error {
	const op = "CheckoutService.Checkout"

	cart, err := s.repoCart.GetOrCreateCart(userID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("خطا در گرفتن سبد خرید")
	}

	items, err := s.repoCart.GetCartItems(cart.ID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("خطا در گرفتن آیتم‌های سبد")
	}

	if len(items) == 0 {
		return richerror.New(op).WithMessage("سبد خرید خالی است")
	}

	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	err = s.repo.Checkout(userID, cart.ID, items, total)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("خطا در Checkout")
	}

	return nil
}
