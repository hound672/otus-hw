//go:build integration

package events

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/tests"
)

func newFakeEvent() *entity.Event {
	event := &entity.Event{}
	_ = gofakeit.Struct(event)

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
}

func (s *eventSuite) TearDownTest() {
}

func TestRunEventSuite(t *testing.T) {
	suite.Run(t, new(eventSuite))
}
