package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/repository/model"
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

func (r *ProductRepository) ShowAll(page int, limit int) ([]model.ProductWithCategory, int, error) {
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
		return []model.ProductWithCategory{}, 0, richerror.New(op).WithErr(err).WithMessage("dont get from dataBase")
	}

	defer rows.Close()

	products := make([]model.ProductWithCategory, 0)

	for rows.Next() {
		var item model.ProductWithCategory

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Stock,
			&item.ImageURL,
			&item.CategoryID,
			&item.CategoryName,
			&item.CategorySlug,
		); err != nil {
			return []model.ProductWithCategory{}, 0, richerror.New(op).WithErr(err).WithMessage("cant scan")
		}

		products = append(products, item)
	}

	if rows.Err() != nil {
		return []model.ProductWithCategory{}, 0, richerror.New(op).WithErr(err).WithMessage("cant scan")
	}

	return products, totoaltems, nil
}

func (r *ProductRepository) ShowByCategory(categorySlug string, page int, limit int) ([]model.ProductWithCategory, int, error) {

	offset := (page - 1) * limit

	const query = `
	SELECT
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.image_url,
		c.id,
		c.name,
		c.slug
	FROM products p
	JOIN categories c ON p.category_id = c.id
	WHERE c.slug = $1
	ORDER BY p.created_at DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(context.Background(), query, categorySlug, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	products := make([]model.ProductWithCategory, 0)

	for rows.Next() {
		var p model.ProductWithCategory
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.ImageURL,
			&p.CategoryID,
			&p.CategoryName,
			&p.CategorySlug,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	var total int
	err = r.db.QueryRow(
		context.Background(),
		`
		SELECT COUNT(*)
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE c.slug = $1
		`,
		categorySlug,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductWithID implements productservice.Repository.
func (r *ProductRepository) GetProductWithID(ctx context.Context, ID int) (entity.Product, error) {

	const op = "postgres.GetProductWithID"
	query := "SELECT id,shop_id, name,description, price, stock, category, image_url,created_at FROM products WHERE id = $1"

	var p entity.Product

	err := r.db.QueryRow(ctx, query, ID).Scan(
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
func (r *ProductRepository) Search(q string, page int) ([]model.ProductWithCategory, int, error) {
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

	var products []model.ProductWithCategory
	for rows.Next() {
		var item model.ProductWithCategory
		if err := rows.Scan(&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Stock,
			&item.ImageURL,
			&item.CategoryID,
			&item.CategoryName,
			&item.CategorySlug); err != nil {
			return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant scan product")
		}
		products = append(products, item)
	}

	if rows.Err() != nil {
		return nil, 0, richerror.New(op).WithErr(rows.Err()).WithMessage("rows error")
	}

	return products, totalItems, nil
}

// DecreaseStock decreases the stock of a product in a transaction-safe way
func (r *ProductRepository) DecreaseStock(ctx context.Context, tx pgx.Tx, productID, qty int) error {
	const op = "postgres.DecreaseStock"

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

func (r *ProductRepository) AddProduct(ctx context.Context, p entity.Product) error {
	const query = `
	INSERT INTO products
	(name, description, price, stock, category_id, image_url)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.CategoryID,
		p.ImageURL,
	)

	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, p entity.Product) error {

	const query = `
	UPDATE products
	SET
		name = $1,
		description = $2,
		price = $3,
		stock = $4,
		category_id = $5,
		image_url = $6,
		updated_at = NOW()
	WHERE id = $7
	`

	cmd, err := r.db.Exec(
		ctx,
		query,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.CategoryID,
		p.ImageURL,
		p.ID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, productID int) error {

	cmd, err := r.db.Exec(
		ctx,
		`DELETE FROM products WHERE id = $1`,
		productID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
