package cartservice

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type Repository interface {
	GetOrCreateCart(userID int) (entity.Cart, error)
	AddItemToCart(cartID, productID, quantity int) error
	UpdateQuantity(cartItemID int, act string) error
	RemoveItem(cartItemID int) error
	GetCartItems(cartID int) ([]entity.CartItem, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
