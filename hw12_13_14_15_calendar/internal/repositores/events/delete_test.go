//go:build integration

package events

import (
	"context"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (s *eventSuite) Test_Delete_Success() {
	ctx := context.Background()

	event := newFakeEvent()

	err := s.repo.Store(ctx, event)
	s.NoError(err)

	err = s.repo.Delete(ctx, event)
	s.NoError(err)

	_, err = s.repo.ReadByUUID(ctx, event.UUID)
	s.ErrorIs(err, entity.ErrEventNotFound)
}
