package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func NewPostgres(cfg Config) (*pgxpool.Pool, error) {
	//str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLMode)
	//
	//db, err := sqlx.Open("postgres", str)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return db, nil

	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return pool, nil
}
