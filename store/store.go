package store

import "github.com/ad-astra-9t/webhook/modelx"

type Store struct {
	WebhookStore
	EventStore
}

func NewStore(modelx *modelx.Modelx) *Store {
	return &Store{
		NewWebhookStore(modelx),
		NewEventStore(modelx),
	}
}
