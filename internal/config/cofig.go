package config

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/repository/postgres"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/authservice"
)

// TODO :  config viber
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	MyPostgres postgres.Config `koanf:"postgres"`
}
