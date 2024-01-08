//go:build integration

package events

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (s *eventSuite) Test_ReadByUUID_Success() {
	ctx := context.Background()

	event := newFakeEvent()

	err := s.repo.Store(ctx, event)
	s.NoError(err)

	eventRead, err := s.repo.ReadByUUID(ctx, event.UUID)
	s.NoError(err)
	s.Equal(event, eventRead)
}

func (s *eventSuite) Test_ReadByUUID_NotFound() {
	ctx := context.Background()

	eventRead, err := s.repo.ReadByUUID(ctx, gofakeit.UUID())
	s.ErrorIs(err, entity.ErrEventNotFound)
	s.Nil(eventRead)
}
