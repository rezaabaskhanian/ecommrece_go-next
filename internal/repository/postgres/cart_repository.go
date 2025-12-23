package postgres

import "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

type CartRepository interface {
	GetOrCreateCart(userID int) (entity.Cart, error)
	AddItemToCart(cartID, productID, quantity int) error
	UpdateQuantity(cartItemID int, act string) error
	RemoveItem(cartItemID int) error
	GetCartItems(cartID int) ([]entity.CartItem, error)
}
