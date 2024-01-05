package events

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) Store(ctx context.Context, event *entity.Event) error {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := sq.Insert(tableName).
		Columns("uuid", "title", "start_at", "end_at", "description", "user_uuid", "notify_before").
		Values(
			event.UUID, event.Title, event.StartAt, event.EndAt, event.Description, event.UserUUID, event.NotifyBefore,
		).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("sq.Insert: %w", err)
	}

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("conn.Exec: %w", err)
	}

	return nil
}
