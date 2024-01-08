package events

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) ReadByUUID(ctx context.Context, uuid string) (*entity.Event, error) {
	conn := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := sq.Select(fields...).
		Where(sq.Eq{"uuid": uuid}).
		From(tableName).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("sq.Select: %w", err)
	}

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("conn.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, entity.ErrEventNotFound
	}

	event := &entity.Event{}
	err = rows.Scan(
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

	return event, nil
}
