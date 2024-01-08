package users

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func New(db *pgxpool.Pool, c *trmpgx.CtxGetter) *Users {
	return &Users{
		db:     db,
		getter: c,
	}
}
