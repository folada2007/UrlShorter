package postgres

import (
	"ShorterAPI/internal/domain/shorter"
	"ShorterAPI/internal/domain/shorter/vo"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type ConnectionPool struct {
	Pool *pgxpool.Pool
}

func InitConnectionPool(connStr string) (*ConnectionPool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &ConnectionPool{pool}, nil
}

func (c ConnectionPool) GetById(id int) shorter.UrlMapping {
	return shorter.UrlMapping{}
}

func (c ConnectionPool) GetByName(name string) shorter.UrlMapping {
	return shorter.UrlMapping{}
}

func (c ConnectionPool) New(vo vo.AliasVO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Pool.Exec(ctx, "INSERT INTO shorter (long_url,short_url) VALUES ($1,$2)", vo.LongUrl, vo.ShortUrl)
	if err != nil {
		return err
	}

	return nil
}
