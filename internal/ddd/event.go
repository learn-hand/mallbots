package ddd

import (
	"time"

	"github.com/google/uuid"
)

type EventOption interface {
	configureEvent(*Event)
}

type Event struct {
	eventId    string
	eventName  string
	payload    any
	metadata   Metadata
	occurredAt time.Time
}

func NewEvent(name string, payload any, options ...EventOption) Event {
	evt := Event{
		eventId:    uuid.New().String(),
		eventName:  name,
		payload:    payload,
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option.configureEvent(&evt)
	}

	return evt
}

func (e Event) EventName() string {
	return e.eventName
}

func (e Event) AggregateId() string {
	return e.metadata.Get(AggregateIdKey).(string)
}

func (e Event) AggregateName() string {
	return e.metadata.Get(AggregateNameKey).(string)
}
