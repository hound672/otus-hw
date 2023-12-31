package events

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Events struct {
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

const (
	tableName = "events"
)

var fields = []string{"uuid", "title", "start_at", "end_at", "description", "user_uuid", "notify_at"}

func New(db *pgxpool.Pool, c *trmpgx.CtxGetter) *Events {
	return &Events{
		db:     db,
		getter: c,
	}
}
