package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

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

func NewFromDatabaseURL() *pgxpool.Pool {
	dbURL := os.Getenv("postgresql://postgres:kMzbeKbcbZFSWhXWoYraKZfVwEbmosiY@crossover.proxy.rlwy.net:39904/railway")
	if dbURL == "" {
		log.Fatal("DATABASE_URL env not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	return pool
}
