package checkoutservcie

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type Repository interface {
	Checkout(UserID, cartID int, items []entity.CartItem, total float64) error
}

type CartRepository interface {
	GetOrCreateCart(userID int) (entity.Cart, error)
	AddItemToCart(cartID, productID, quantity int) error
	UpdateQuantity(cartItemID int, act string) error
	RemoveItem(cartItemID int) error
	GetCartItems(cartID int) ([]entity.CartItem, error)
}

type Service struct {
	repo     Repository
	repoCart CartRepository
}

func New(repo Repository, repoCart CartRepository) Service {
	return Service{repo: repo, repoCart: repoCart}

}
