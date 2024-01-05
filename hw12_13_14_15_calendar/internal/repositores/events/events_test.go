//go:build integration

package events

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/tests"
)

func newFakeEvent() *entity.Event {
	event := &entity.Event{}
	_ = gofakeit.Struct(event)
	event.UUID = gofakeit.UUID()
	event.UserUUID = gofakeit.UUID()
	event.NotifyBefore = gofakeit.UintRange(0, 0xFFFFFFFF)
	// truncate nsecs
	event.StartAt = gofakeit.FutureDate().UTC().Truncate(time.Second)
	event.EndAt = gofakeit.FutureDate().UTC().Truncate(time.Second)

	return event
}

type eventSuite struct {
	suite.Suite

	psqlTests *tests.PsqlTests
	repo      *Events
}

func (s *eventSuite) SetupSuite() {
	// set up container with postgres

	s.psqlTests = tests.GetPsql()
	s.repo = New(s.psqlTests.Pool, s.psqlTests.CtxGetter)
}

func (s *eventSuite) TearDownSuite() {
	s.psqlTests.Cleanup()
	err := s.psqlTests.Container.Terminate(context.Background())
	if err != nil {
		panic(err)
	}
}

func (s *eventSuite) SetupTest() {
	_, err := s.psqlTests.Migration.Up()
	if err != nil {
		panic(err)
	}
}

func (s *eventSuite) TearDownTest() {
	_, err := s.psqlTests.Migration.Down()
	if err != nil {
		panic(err)
	}
}

func TestRunEventSuite(t *testing.T) {
	suite.Run(t, new(eventSuite))
}
