package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/model"
)

type Store interface {
	CreateWebhook(target domain.Webhook) error
	GetWebhook(target domain.Webhook) (result domain.Webhook, err error)
}

type DefaultStore struct {
	*model.DefaultModel
	WebhookStore
	EventStore
}

func NewDefaultStore(model *model.DefaultModel, adapt *model.ModelAdapt) *DefaultStore {
	return &DefaultStore{
		model,
		NewWebhookStore(model, adapt),
		NewEventStore(model),
	}
}
