package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/modelx"
)

type WebhookStore struct {
	modelx     *modelx.Modelx
	modeladapt domain.Adapt[domain.Webhook, modelx.Webhook]
}

func (s WebhookStore) CreateWebhook(target domain.Webhook) error {
	modeltarget := s.modeladapt.AdaptTarget(target)
	return s.modelx.CreateWebhook(modeltarget)
}

func (s WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.modeladapt.AdaptTarget(target)

	modelresult, err := s.modelx.GetWebhook(modeltarget)
	if err != nil {
		return result, err
	}

	result = s.modeladapt.AdaptDomain(modelresult)

	return result, err
}

func NewWebhookStore(modelx *modelx.Modelx, adapt *modelx.ModelAdapt) WebhookStore {
	return WebhookStore{
		modelx,
		adapt,
	}
}
