package repository

import "github.com/jackc/pgx/v5/pgxpool"

type DBRepository struct {
	pool *pgxpool.Pool
}
