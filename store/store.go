package store

import "github.com/ad-astra-9t/webhook/modelx"

type Store struct {
	WebhookStore
}

func NewStore(model *modelx.Model) *Store {
	return &Store{
		WebhookStore: NewWebhookStore(model),
	}
}
