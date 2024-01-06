package events

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) Store(ctx context.Context, event *entity.Event) error {
	onConfictQuery := `
	ON CONFLICT (uuid) DO UPDATE 
	SET title=EXCLUDED.title,
		start_at=EXCLUDED.start_at,
		end_at=EXCLUDED.end_at,
		description=EXCLUDED.description,
		user_uuid=EXCLUDED.user_uuid,
		notify_at=EXCLUDED.notify_at
`

	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := sq.Insert(tableName).
		Columns("uuid", "title", "start_at", "end_at", "description", "user_uuid", "notify_at").
		Values(
			event.UUID, event.Title, event.StartAt, event.EndAt, event.Description, event.UserUUID, event.NotifyAt,
		).Suffix(onConfictQuery).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("sq.Insert: %w", err)
	}

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("conn.Exec: %w", err)
	}

	return nil
}
