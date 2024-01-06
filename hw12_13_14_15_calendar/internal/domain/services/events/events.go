package events

import (
	"context"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

type (
	Repo interface {
		Store(ctx context.Context, event *entity.Event) error
		GetAll(ctx context.Context) ([]*entity.Event, error)
		Delete(ctx context.Context, event *entity.Event) error
	}

	Events struct {
		eventsRepo Repo
	}
)

func New(eventsRepo Repo) *Events {
	events := &Events{
		eventsRepo: eventsRepo,
	}
	return events
}
