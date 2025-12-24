package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
)

type OrderRepository struct {
	DB *pgxpool.Pool
}

// -------------------------------
// Constructor
// -------------------------------

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{DB: db}
}

// BeginTx implements checkoutservcie.BeginRepository.
func (r *OrderRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	panic("unimplemented")
}

func (r *OrderRepository) CreateOrder(ctx context.Context,
	tx pgx.Tx, order entity.Order) (int, error) {
	const query = `
	INSERT INTO orders (user_id, cart_id, status, total)
	VALUES ($1,$2,$3,$4)
	RETURNING id
	`
	var id int
	err := r.DB.QueryRow(
		context.Background(),
		query,
		order.USerID,
		order.CartID,
		order.Status,
		order.Total,
	).Scan(&id)

	return id, err
}

func (r *OrderRepository) CreateOrderItems(ctx context.Context,
	tx pgx.Tx, items []entity.OrderItem) error {
	const query = `
	INSERT INTO order_items (order_id, product_id, name, price, quantity)
	VALUES ($1,$2,$3,$4,$5)
	`

	for _, it := range items {
		_, err := r.DB.Exec(
			context.Background(),
			query,
			it.OrderID,
			it.ProductID,
			it.Name,
			it.Price,
			it.Quantity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
