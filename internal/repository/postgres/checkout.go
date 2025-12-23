package postgres

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (r *Repository) Checkout(userID, cartID int, items []entity.CartItem, total float64) error {

	const op = "postgres.Checkout"

	tx, err := r.DB.Begin(context.Background())

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		} else {
			_ = tx.Commit(context.Background())
		}
	}()

	var orderID int
	err = tx.QueryRow(context.Background(),
		`INSERT INTO orders(user_id, cart_id, status, total) VALUES($1,$2,$3,$4) RETURNING id`,
		userID, cartID, "pending_payment", total,
	).Scan(&orderID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("خطا در ایجاد سفارش")
	}

	for _, item := range items {
		// کاهش stock
		res, err := tx.Exec(context.Background(),
			`UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`,
			item.Quantity, item.ProductID,
		)
		if err != nil {
			return richerror.New(op).WithErr(err).WithMessage("خطا در کاهش موجودی")
		}

		if res.RowsAffected() == 0 {
			return richerror.New(op).WithMessage("موجودی کافی نیست")
		}

		_, err = tx.Exec(context.Background(),
			`INSERT INTO order_items(order_id, product_id, name, price, quantity) VALUES($1,$2,$3,$4,$5)`,
			orderID, item.ProductID, item.Name, item.Price, item.Quantity,
		)
		if err != nil {
			return richerror.New(op).WithErr(err).WithMessage("خطا در ایجاد OrderItem")
		}
	}

	_, err = tx.Exec(context.Background(), `UPDATE cart SET status='checkout' WHERE id=$1`, cartID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage("خطا در تغییر وضعیت cart")
	}

	return nil

}
