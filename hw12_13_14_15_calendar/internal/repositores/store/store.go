package store

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func New(db *pgxpool.Pool, c *trmpgx.CtxGetter) *Store {
	return &Store{
		db:     db,
		getter: c,
	}
}
