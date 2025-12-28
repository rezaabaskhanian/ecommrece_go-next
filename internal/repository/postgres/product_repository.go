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
type ProductRepository struct {
	db *pgxpool.Pool
}

// -------------------------------
// Constructor
// -------------------------------
func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// -------------------------------
// Methods
// -------------------------------

func (r *ProductRepository) ShowAll(page int, limit int) ([]entity.Product, int, error) {
	const op = "postgres.ShowAll"

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	var totoaltems int

	countQuery := `SELECT COUNT(*) FROM products`

	if err := r.db.QueryRow(context.Background(), countQuery).Scan(&totoaltems); err != nil {
		return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant count products")

	}

	query := "SELECT id,shop_id, name,description, price, stock, category, image_url  FROM products ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(context.Background(), query, limit, offset)

	if err != nil {
		return []entity.Product{}, 0, richerror.New(op).WithErr(err).WithMessage("dont get from dataBase")
	}

	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var item entity.Product

		if err := rows.Scan(
			&item.ID,
			&item.ShopID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Stock,
			&item.CategoryID,
			&item.ImageURL,
		); err != nil {
			return []entity.Product{}, 0, richerror.New(op).WithErr(err).WithMessage("cant scan")
		}

		products = append(products, item)
	}

	if rows.Err() != nil {
		return []entity.Product{}, 0, richerror.New(op).WithErr(err).WithMessage("cant scan")
	}

	return products, totoaltems, nil
}

// GetProductWithID implements productservice.Repository.
func (r *ProductRepository) GetProductWithID(ID int) (entity.Product, error) {

	const op = "postgres.GetProductWithID"
	query := "SELECT id,shop_id, name,description, price, stock, category, image_url,created_at FROM products WHERE id = $1"

	var p entity.Product

	err := r.db.QueryRow(context.Background(), query, ID).Scan(
		&p.ID, &p.ShopID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.ImageURL, &p.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Product{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		return entity.Product{}, err
	}

	return p, nil

}

// Search returns products that match the query string
func (r *ProductRepository) Search(q string, page int) ([]entity.Product, int, error) {
	const op = "postgres.ProductRepository.Search"
	if page <= 0 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit
	searchValue := "%" + q + "%"

	var totalItems int
	countQuery := `
		SELECT COUNT(*) FROM products
		WHERE name ILIKE $1 OR description ILIKE $1 OR category ILIKE $1
	`
	if err := r.db.QueryRow(context.Background(), countQuery, searchValue).Scan(&totalItems); err != nil {
		return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant count search products")
	}

	query := `
		SELECT id, shop_id, name, description, price, stock, category, image_url, created_at
		FROM products
		WHERE name ILIKE $1 OR description ILIKE $1 OR category ILIKE $1
		ORDER BY id
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(context.Background(), query, searchValue, limit, offset)
	if err != nil {
		return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant search products")
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(&p.ID, &p.ShopID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.ImageURL, &p.CreatedAt); err != nil {
			return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant scan product")
		}
		products = append(products, p)
	}

	if rows.Err() != nil {
		return nil, 0, richerror.New(op).WithErr(rows.Err()).WithMessage("rows error")
	}

	return products, totalItems, nil
}

// DecreaseStock decreases the stock of a product in a transaction-safe way
func (r *ProductRepository) DecreaseStock(ctx context.Context, tx pgx.Tx, productID, qty int) error {
	const op = "postgres.ProductRepository.DecreaseStock"

	res, err := tx.Exec(ctx,
		`UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`,
		qty,
		productID,
	)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}
	if res.RowsAffected() == 0 {
		return richerror.New(op).WithMessage("insufficient stock")
	}
	return nil
}
