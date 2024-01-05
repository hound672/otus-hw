package events

import (
	"context"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/domain/entity"
)

func (r *Events) Delete(_ context.Context, _ *entity.Event) error {
	return nil
}
