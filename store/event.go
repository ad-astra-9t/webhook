package store

import "github.com/ad-astra-9t/webhook/modelx"

type EventStore struct {
	model modelx.Model
}

func NewEventStore(model modelx.Model) EventStore {
	return EventStore{model}
}
