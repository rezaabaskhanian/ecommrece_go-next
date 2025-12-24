package checkoutservcie

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
)

type CartRepository interface {
	GetOrCreateCart(userID int) (entity.Cart, error)
	GetCartItems(cartID int) ([]entity.CartItem, error)
	ClearCart(ctx context.Context, tx pgx.Tx, cartID int) error
}

type ProductRepository interface {
	GetProductWithID(ID int) (entity.Product, error)
	DecreaseStock(ctx context.Context, tx pgx.Tx, productID int, qty int) error
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, tx pgx.Tx, order entity.Order) (int, error)
	CreateOrderItems(ctx context.Context, tx pgx.Tx, items []entity.OrderItem) error
}

type BeginRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

type Service struct {
	cartRepo    CartRepository
	productRepo ProductRepository
	orderRepo   OrderRepository

	beginRepo BeginRepository
}

func New(
	cartRepo CartRepository,
	productRepo ProductRepository,
	orderRepo OrderRepository,
	beginRepo BeginRepository,
) Service {
	return Service{
		cartRepo:    cartRepo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
		beginRepo:   beginRepo,
	}
}
