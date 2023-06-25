package ddd

import (
	"context"
	"sync"
)

type EventHandler func(ctx context.Context, event Event) error

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: map[string][]EventHandler{},
	}
}

func (d *EventDispatcher) Subcribe(event Event, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[event.EventName()] = append(d.handlers[event.EventName()], handler)
}

func (d *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range d.handlers[event.EventName()] {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
