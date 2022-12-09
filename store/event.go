package store

import "github.com/ad-astra-9t/webhook/modelx"

type EventStore struct {
	modelx *modelx.Modelx
}

func NewEventStore(modelx *modelx.Modelx) EventStore {
	return EventStore{modelx}
}
