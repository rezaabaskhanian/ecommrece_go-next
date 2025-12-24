package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

// -------------------------------
// Struct
// -------------------------------
type CartRepository struct {
	DB *pgxpool.Pool
}

// -------------------------------
// Constructor
// -------------------------------

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{DB: db}
}

// func (r *Repository) CreateCart(userID int) (entity.Cart, error) {
// 	const op = "postgres.CreateCart"
// 	query := `INSERT INTO cart (user_id, status)VALUES($1,$2)RETURNING id, status, created_at, updated_at`

// 	var c entity.Cart
// 	Status := "active"

// 	if err := r.DB.QueryRow(context.Background(), query, userID, Status).Scan(&c.ID, &c.Status, &c.CreatedAt, &c.UpdatedAt); err != nil {
// 		return entity.Cart{}, richerror.New(op).WithErr(err)

// 	}

// 	return entity.Cart{
// 		ID:        c.ID,
// 		UserID:    userID,
// 		Status:    c.Status,
// 		CreatedAt: c.CreatedAt,
// 		UpdatedAt: c.UpdatedAt,
// 	}, nil

// }
// func (r *Repository) CreateCartItem(cartID, productID, quantity int) (entity.CartItem, error) {
// 	const op = "postgres.CreateCartItem"

// 	var product struct {
// 		Name     string
// 		Price    float64
// 		ImageURL string
// 		Stock    int // برای چک موجودی
// 	}

// 	productQuery := `SELECT name, price, image_url ,stock FROM products WHERE id = $1`
// 	err := r.DB.QueryRow(context.Background(), productQuery, productID).Scan(
// 		&product.Name,
// 		&product.Price,
// 		&product.ImageURL,
// 		&product.Stock,
// 	)

// 	if err == pgx.ErrNoRows {
// 		return entity.CartItem{}, richerror.New(op).
// 			WithMessage("محصول یافت نشد")
// 	}

// 	if err != nil {
// 		return entity.CartItem{}, richerror.New(op).WithErr(err).WithMessage("محصول یافت نشد")
// 	}

// 	if product.Stock < quantity {
// 		return entity.CartItem{}, richerror.New(op).
// 			WithMessage("موجودی کافی نیست")
// 	}

// 	var existingItemID int
// 	var existingQuantity int

// 	checkQuery := `
// 		SELECT cart_item_id, quantity
// 		FROM cart_item
// 		WHERE cart_id = $1 AND product_id = $2`

// 	err = r.DB.QueryRow(context.Background(), checkQuery, cartID, productID).Scan(
// 		&existingItemID,
// 		&existingQuantity,
// 	)

// 	// if existing and update

// 	if err == nil {
// 		updateQuery := `
// 			UPDATE cart_item
// 			SET quantity = $1
// 			WHERE cart_item_id = $2
// 			RETURNING cart_item_id,  product_id, quantity, name, price, image_url
// 		`
// 		var item entity.CartItem
// 		err = r.DB.QueryRow(context.Background(), updateQuery, existingQuantity+quantity, existingItemID).Scan(
// 			&item.CartItemID,
// 			&item.ProductID,
// 			&item.Quantity,
// 			&item.Name,
// 			&item.Price,
// 			&item.ImageURL,
// 		)

// 		if err != nil {
// 			return entity.CartItem{}, richerror.New(op).WithErr(err)
// 		}

// 		return item, nil
// 	}

// 	// if not existing

// 	insertQuery := `INSERT INTO cart_item(cart_id,product_id,quantity,name,price,image_url)VALUES($1,$2,$3,$4,$5,$6)RETURNING cart_item_id `

// 	var c entity.CartItem

// 	if err := r.DB.QueryRow(context.Background(), insertQuery, cartID, productID, quantity, product.Name, product.Price, product.ImageURL).
// 		Scan(&c.CartItemID, &c.ProductID, &c.Quantity, &c.Name, &c.Price, &c.ImageURL); err != nil {
// 		return entity.CartItem{}, richerror.New(op).WithErr(err).WithMessage("خطا در اضافه کردن محصول")

// 	}

// 	return c, nil

// }

// func (r *Repository) UpdateQuantity(cartItemID, productID int, act string) (entity.CartItem, error) {
// 	const op = "postgres.UpdateQuantity"
// 	var query string
// 	if act == "add" {
// 		query = `
// 		UPDATE cart_item
// 		SET quantity = quantity + 1
// 		WHERE id = $1 AND product_id = $2
// 		RETURNING id, cart_id, product_id, quantity
// 	`

// 	} else {
// 		query = `
// 		UPDATE cart_item
// 		SET quantity = quantity - 1
// 		WHERE id = $1 AND product_id = $2
// 		RETURNING id, cart_id, product_id, quantity
// 	`

// 	}

// 	var item entity.CartItem

// 	err := r.DB.QueryRow(context.Background(), query, cartItemID, productID).Scan(

// 		&item.CartItemID, &item.ProductID, &item.Quantity, &item.Name, &item.Price, &item.ImageURL,
// 	)

// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return entity.CartItem{}, richerror.New(op).
// 				WithMessage("cart item not found")
// 		}
// 		return entity.CartItem{}, richerror.New(op).WithErr(err)
// 	}

// 	return item, nil

// }

func (r *CartRepository) GetCart(userID int) (entity.Cart, error) {
	const op = "postgres.GetCart"

	query := `SELECT id, user_id, status, created_at, updated_at FROM cart WHERE user_id = $1`

	var c entity.Cart

	if err := r.DB.QueryRow(context.Background(), query, userID).Scan(
		&c.ID, &c.UserID, &c.Status, c.CreatedAt, c.UpdatedAt,
	); err != nil {

		if err == pgx.ErrNoRows {
			return entity.Cart{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		return entity.Cart{}, richerror.New(op).WithErr(err).WithMessage("dont get carts")
	}

	return c, nil

}

// func (r *Repository) GetCartItem(cartID int) ([]entity.CartItem, error) {

// 	const op = "postgres.GetCartItems"

// 	query := `
// 			SELECT
// 				ci.id AS cart_item_id,
// 				ci.cart_id,
// 				ci.quantity,
// 				p.id AS product_id,
// 				p.name,
// 				p.price,
// 				p.image_url
// 			FROM cart_item ci
// 			JOIN products p ON p.id = ci.product_id
// 			WHERE ci.cart_id = $1
// 			ORDER BY ci.id
// 		`

// 	rows, err := r.DB.Query(context.Background(), query, cartID)
// 	if err != nil {
// 		return nil, richerror.New(op).
// 			WithErr(err).
// 			WithMessage("خطا در دریافت آیتم‌های سبد").
// 			WithKind(richerror.KindUnexpected)
// 	}
// 	defer rows.Close()

// 	var cartItems []entity.CartItem

// 	for rows.Next() {
// 		var item entity.CartItem

// 		err := rows.Scan(
// 			&item.CartItemID, // ci.id

// 			&item.Quantity,  // ci.quantity
// 			&item.ProductID, // p.id
// 			&item.Name,      // p.name
// 			&item.Price,     // p.price
// 			&item.ImageURL,  // p.image_url
// 		)

// 		if err != nil {
// 			return nil, richerror.New(op).
// 				WithErr(err).
// 				WithMessage("خطا در خواندن آیتم سبد").
// 				WithKind(richerror.KindUnexpected)
// 		}

// 		cartItems = append(cartItems, item)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, richerror.New(op).
// 			WithErr(err).
// 			WithMessage("خطا در پردازش آیتم‌ها").
// 			WithKind(richerror.KindUnexpected)
// 	}

// 	return cartItems, nil
// }

// func (r *Repository) GetCartWithItems(userID int) (entity.CartWithItem, error) {
// 	const op = "postgres.GetCartWithItems"

// 	cart, err := r.GetCart(userID)
// 	if err != nil {
// 		return entity.CartWithItem{}, richerror.New(op).WithErr(err).WithMessage("dont get from cart")
// 	}

// 	items, err := r.GetCartItem(cart.ID)
// 	if err != nil {
// 		return entity.CartWithItem{}, richerror.New(op).WithErr(err).WithMessage("dont get from cartItem")
// 	}

// 	totalPrice := 0.0
// 	for _, item := range items {
// 		totalPrice += item.Price * float64(item.Quantity)
// 	}

// 	return entity.CartWithItem{
// 		Cart:       cart,
// 		Items:      items,
// 		TotalPrice: totalPrice,
// 		ItemCount:  len(items),
// 	}, nil
// }

func (r *CartRepository) GetOrCreateCart(userID int) (entity.Cart, error) {
	const op = "postgres.GetOrCreateCart"

	query := `
	SELECT id, user_id, status, created_at, updated_at
	FROM cart
	WHERE user_id = $1 AND status = 'active'
	`

	var c entity.Cart
	err := r.DB.QueryRow(context.Background(), query, userID).
		Scan(&c.ID, &c.UserID, &c.Status, &c.CreatedAt, &c.UpdatedAt)

	if err == nil {
		return c, nil
	}

	if err != pgx.ErrNoRows {
		return entity.Cart{}, richerror.New(op).WithErr(err)
	}

	insert := `
	INSERT INTO cart (user_id, status)
	VALUES ($1, 'active')
	RETURNING id, user_id, status, created_at, updated_at
	`

	err = r.DB.QueryRow(context.Background(), insert, userID).
		Scan(&c.ID, &c.UserID, &c.Status, &c.CreatedAt, &c.UpdatedAt)

	if err != nil {
		return entity.Cart{}, richerror.New(op).WithErr(err)
	}

	return c, nil
}

func (r *CartRepository) AddItemToCart(cartID, productID, quantity int) error {
	const op = "postgres.AddItemToCart"

	query := `
	INSERT INTO cart_item (cart_id, product_id, quantity)
	VALUES ($1, $2, $3)
	ON CONFLICT (cart_id, product_id)
	DO UPDATE SET quantity = cart_item.quantity + EXCLUDED.quantity
	`

	_, err := r.DB.Exec(context.Background(), query, cartID, productID, quantity)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return nil
}

func (r *CartRepository) UpdateQuantity(cartItemID int, act string) error {
	const op = "postgres.UpdateQuantity"

	var query string
	if act == "add" {
		query = `
		UPDATE cart_item
		SET quantity = quantity + 1
		WHERE id = $1
		`
	} else {
		query = `
		UPDATE cart_item
		SET quantity = quantity - 1
		WHERE id = $1 AND quantity > 1
		`
	}

	cmd, err := r.DB.Exec(context.Background(), query, cartItemID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if cmd.RowsAffected() == 0 {
		return richerror.New(op).WithMessage("cart item not found")
	}

	return nil
}

func (r *CartRepository) GetCartItems(cartID int) ([]entity.CartItem, error) {
	const op = "postgres.GetCartItems"

	query := `
	SELECT
		ci.id,
		ci.product_id,
		ci.quantity,
		p.name,
		p.price,
		p.image_url
	FROM cart_item ci
	JOIN products p ON p.id = ci.product_id
	WHERE ci.cart_id = $1
	ORDER BY ci.id
	`

	rows, err := r.DB.Query(context.Background(), query, cartID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}
	defer rows.Close()

	var items []entity.CartItem

	for rows.Next() {
		var item entity.CartItem
		err := rows.Scan(
			&item.CartItemID,
			&item.ProductID,
			&item.Quantity,
			&item.Name,
			&item.Price,
			&item.ImageURL,
		)
		if err != nil {
			return nil, richerror.New(op).WithErr(err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *CartRepository) GetCartWithItems(userID int) (entity.CartWithItem, error) {
	const op = "postgres.GetCartWithItems"

	cart, err := r.GetOrCreateCart(userID)
	if err != nil {
		return entity.CartWithItem{}, richerror.New(op).WithErr(err)
	}

	items, err := r.GetCartItems(cart.ID)
	if err != nil {
		return entity.CartWithItem{}, richerror.New(op).WithErr(err)
	}

	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return entity.CartWithItem{
		Cart:       cart,
		Items:      items,
		TotalPrice: total,
		ItemCount:  len(items),
	}, nil
}

func (r *CartRepository) RemoveItem(cartItemID int) error {
	const op = "postgres.RemoveItem"

	cmd, err := r.DB.Exec(
		context.Background(),
		`DELETE FROM cart_item WHERE id = $1`,
		cartItemID,
	)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if cmd.RowsAffected() == 0 {
		return richerror.New(op).WithMessage("item not found")
	}

	return nil
}

func (r *CartRepository) ClearCart(ctx context.Context,
	tx pgx.Tx, cartID int) error {
	const op = "postgres.Clear"
	_, err := r.DB.Exec(
		context.Background(),
		`DELETE FROM cart_items WHERE cart_id=$1`,
		cartID,
	)

	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	return err
}
