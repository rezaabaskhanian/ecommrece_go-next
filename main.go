package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
)

func main() {
	// اتصال به دیتابیس
	connStr := "postgres://user:pass@localhost:5433/ecommerce?sslmode=disable"
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	for i := 1; i <= 30; i++ {
		product := entity.Product{
			ShopID:      1,
			Name:        fmt.Sprintf("Terrarium %d", i),
			Description: fmt.Sprintf("This is Terrarium number %d", i),
			Price:       int64(50 + i), // قیمت تصادفی ساده
			Stock:       10 + i%5,      // موجودی تصادفی
			CategoryID:  5,
			ImageURL:    "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		query := `
			INSERT INTO products
			(shop_id, name, description, price, stock, category, image_url, created_at, updated_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`

		_, err := db.Exec(context.Background(),
			query,
			product.ShopID,
			product.Name,
			product.Description,
			product.Price,
			product.Stock,
			product.CategoryID,
			product.ImageURL,
			product.CreatedAt,
			product.UpdatedAt,
		)
		if err != nil {
			log.Println("Failed to insert product:", err)
		} else {
			fmt.Println("Inserted:", product.Name)
		}
	}
}
