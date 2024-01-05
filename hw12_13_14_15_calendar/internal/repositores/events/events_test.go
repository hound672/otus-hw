//go:build integration

package events

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
)

func newFakeEvent() *entity.Event {
	event := &entity.Event{}
	_ = gofakeit.Struct(event)

	return event
}

type eventSuite struct {
	suite.Suite

	repo *Events
}

func (s *eventSuite) SetupSuite() {
	// set up container with postgres
	logger.Info("SETUP SUITE")
}

func (s *eventSuite) SetupTest() {
	logger.Info("SETUP TEST")
}

func (s *eventSuite) TearDownTest() {
	logger.Info("TEAR DOWN TEST")
}

func TestRunEventSuite(t *testing.T) {
	suite.Run(t, new(eventSuite))
}
