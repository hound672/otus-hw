//go:build integration

package events

import (
	"context"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (s *eventSuite) Test_GetAll_Success() {
	ctx := context.Background()

	events := make([]*entity.Event, 10)
	for i := range events {
		events[i] = newFakeEvent()
		err := s.repo.Store(ctx, events[i])
		s.NoError(err)
	}

	eventsRead, err := s.repo.ReadAll(ctx)
	s.NoError(err)
	s.ElementsMatch(events, eventsRead)
}
