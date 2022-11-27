package store

import "github.com/ad-astra-9t/webhook/model"

type Store struct {
	WebhookStore
}

func NewStore(model *model.Model) *Store {
	return &Store{
		WebhookStore: NewWebhookStore(model),
	}
}
