package entity

import (
	"time"
)

type Event struct {
	UUID  string
	Title string
	// StartAt time of event starting
	StartAt time.Time
	// EndAt time of event finished
	EndAt        time.Time
	Description  string
	UserUUID     string
	NotifyBefore time.Duration
}

func NewEvent(
	uuid string,
	title string,
	startAt time.Time,
	endAt time.Time,
	description string,
	userUUID string,
	notifyBefore time.Duration,
) *Event {
	return &Event{
		UUID:         uuid,
		Title:        title,
		StartAt:      startAt,
		EndAt:        endAt,
		Description:  description,
		UserUUID:     userUUID,
		NotifyBefore: notifyBefore,
	}
}
