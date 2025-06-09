package postgres

import (
	"ShorterAPI/internal/domain/shorter"
	"ShorterAPI/internal/domain/shorter/vo"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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

func (c ConnectionPool) New(vo vo.AliasVO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Pool.Exec(ctx, "INSERT INTO shorter (long_url,short_url) VALUES ($1,$2)", vo.LongUrl, vo.ShortUrl)

	if err != nil {
		return err
	}

	return nil
}

func (c ConnectionPool) FindLongUrlByKey(shortUrl string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var longUrl string

	err := c.Pool.QueryRow(ctx, "SELECT long_url FROM shorter WHERE short_url=$1", shortUrl).Scan(&longUrl)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", shorter.ErrNotFound
	}

	if err != nil {
		return "", err
	}
	return longUrl, nil
}
