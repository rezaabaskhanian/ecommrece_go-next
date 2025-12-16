package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

// beacase declare in user.go now comment

// type Repository struct {
// 	DB *pgxpool.Pool
// }

func MyNewPostgresProduct(db *pgxpool.Pool) ProductRepository {
	return &Repository{DB: db}

}

// GetProductWithID implements ProductRepository.
func (r *Repository) GetProductWithID(ID int) (entity.Product, error) {
	const op = "postgres.GetProductWithID"
	query := "SELECT id,shop_id, name,description, price, stock, category, image_url,created_at FROM products WHERE id = $1"

	var p entity.Product

	err := r.DB.QueryRow(context.Background(), query, ID).Scan(
		&p.ID, &p.ShopID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.ImageURL, &p.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Product{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		return entity.Product{}, err
	}

	return p, nil

}

// ShowAll implements ProductRepository.
func (r *Repository) ShowAll(page, limit int) ([]entity.Product, int, error) {

	const op = "postgres.ShowAll"

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	var totoaltems int

	countQuery := `SELECT COUNT(*) FROM products`

	if err := r.DB.QueryRow(context.Background(), countQuery).Scan(&totoaltems); err != nil {
		return nil, 0, richerror.New(op).WithErr(err).WithMessage("cant count products")

	}

	query := "SELECT id,shop_id, name,description, price, stock, category, image_url  FROM products ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := r.DB.Query(context.Background(), query, limit, offset)

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
			&item.Category,
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
