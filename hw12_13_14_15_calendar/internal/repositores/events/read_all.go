package events

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) ReadAll(ctx context.Context) ([]*entity.Event, error) {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := sq.Select(fields...).From(tableName).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("sq.Select: %w", err)
	}

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("conn.Query: %w", err)
	}
	defer rows.Close()

	events := make([]*entity.Event, 0)
	for rows.Next() {
		event := &entity.Event{}
		err := rows.Scan(
			&event.UUID,
			&event.Title,
			&event.StartAt,
			&event.EndAt,
			&event.Description,
			&event.UserUUID,
			&event.NotifyAt,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		events = append(events, event)
	}

	return events, nil
}
