package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	UserName string
	Password string
	Port     int
	Host     string
	DBName   string
}

func New(cfg Config) *pgxpool.Pool {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	return pool
}
