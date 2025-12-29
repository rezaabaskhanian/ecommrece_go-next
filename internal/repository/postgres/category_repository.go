package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

type CategoryRepository struct {
	DB *pgxpool.Pool
}

// -------------------------------
// Constructor
// -------------------------------

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetList(ctx context.Context, limit int) ([]entity.Category, error) {
	const op = "postgres.GetList"

	rows, err := r.DB.Query(ctx,
		`SELECT id, name, slug FROM categories ORDER BY id ASC LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var c entity.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
			return nil, richerror.New(op).WithErr(err)
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByName(ctx context.Context, name string) (entity.Category, error) {
	const op = "postgres.GetCategoryWithID"

	query := "SELECT id, name, slug FROM categories WHERE name = $1"

	var c entity.Category

	err := r.DB.QueryRow(ctx, query, name).Scan(
		&c.ID, &c.Name, &c.Slug)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Category{}, richerror.
				New(op).
				WithKind(richerror.KindNotFound)
		}

		return entity.Category{}, richerror.
			New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return c, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, Id int) (entity.Category, error) {
	const op = "postgres.GetCategoryWithID"

	query := "SELECT id, name, slug FROM categories WHERE id = $1"

	var c entity.Category

	err := r.DB.QueryRow(ctx, query, Id).Scan(
		&c.ID, &c.Name, &c.Slug)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Category{}, richerror.
				New(op).
				WithKind(richerror.KindNotFound)
		}

		return entity.Category{}, richerror.
			New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return c, nil
}

func (r *CategoryRepository) AddCategory(ctx context.Context,
	cat entity.Category) error {
	const op = "postgres.category.AddCategory"

	const query = `
	INSERT INTO categories (name, slug)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		cat.Name,
		cat.Slug,
	)

	if err != nil {
		return richerror.
			New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, cat entity.Category) error {

	const op = "postgres.category.Update"

	const query = `
		UPDATE categories
		SET name = $1,
		    slug = $2,
		    updated_at = NOW()
		WHERE id = $3
	`

	cmd, err := r.DB.Exec(
		ctx,
		query,
		cat.Name,
		cat.Slug,
		cat.ID,
	)

	if err != nil {
		return richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	if cmd.RowsAffected() == 0 {
		return richerror.New(op).
			WithKind(richerror.KindNotFound)
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(
	ctx context.Context,
	id int,
) error {

	const op = "postgres.category.Delete"

	const query = `
		DELETE FROM categories
		WHERE id = $1
	`

	cmd, err := r.DB.Exec(ctx, query, id)
	if err != nil {
		return richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	if cmd.RowsAffected() == 0 {
		return richerror.New(op).
			WithKind(richerror.KindNotFound)
	}

	return nil
}
