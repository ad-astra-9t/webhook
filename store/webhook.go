package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/modelx"
)

type WebhookStore struct {
	model      modelx.Model
	modeladapt domain.Adapt[domain.Webhook, modelx.Webhook]
}

func (s WebhookStore) CreateWebhook(target domain.Webhook) error {
	modeltarget := s.modeladapt.AdaptTarget(target)
	return s.model.CreateWebhook(modeltarget)
}

func (s WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.modeladapt.AdaptTarget(target)

	modelresult, err := s.model.GetWebhook(modeltarget)
	if err != nil {
		return result, err
	}

	result = s.modeladapt.AdaptDomain(modelresult)

	return result, err
}

func NewWebhookStore(model modelx.Model, adapt *modelx.ModelAdapt) WebhookStore {
	return WebhookStore{
		model,
		adapt,
	}
}
