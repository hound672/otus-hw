package events

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) Delete(ctx context.Context, event *entity.Event) error {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := sq.Delete(tableName).Where(sq.Eq{"uuid": event.UUID}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("sq.Delete: %w", err)
	}

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("conn.Exec: %w", err)
	}

	return nil
}
