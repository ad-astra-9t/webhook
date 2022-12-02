package store

import "github.com/ad-astra-9t/webhook/model"

type EventStore struct {
	model *model.Model
}

func NewEventStore(model *model.Model) EventStore {
	return EventStore{
		model: model,
	}
}
