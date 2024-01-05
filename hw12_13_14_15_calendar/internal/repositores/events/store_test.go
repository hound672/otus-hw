//go:build integration

package events

import (
	"context"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (s *eventSuite) Test_Store_Success() {
	ctx := context.Background()

	event := newFakeEvent()

	err := s.repo.Store(ctx, event)
	s.NoError(err)

	rows, err := s.psqlTests.Pool.Query(
		ctx,
		"SELECT uuid, title, start_at, end_at, description, user_uuid, notify_before FROM events WHERE uuid = $1",
		event.UUID,
	)
	s.NoError(err)
	defer rows.Close()

	eventRead := &entity.Event{}
	for rows.Next() {
		err = rows.Scan(
			&eventRead.UUID,
			&eventRead.Title,
			&eventRead.StartAt,
			&eventRead.EndAt,
			&eventRead.Description,
			&eventRead.UserUUID,
			&eventRead.NotifyBefore,
		)
		s.NoError(err)
	}

	s.Equal(event, eventRead)
}
