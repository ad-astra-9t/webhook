package store

import "github.com/ad-astra-9t/webhook/modelx"

type Store struct {
	WebhookStore
}

func NewStore(modelx *modelx.Modelx) *Store {
	return &Store{
		WebhookStore: NewWebhookStore(modelx),
	}
}
