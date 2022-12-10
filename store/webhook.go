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

	txmodel, err := s.modelx.AutoTx()
	if err != nil {
		return err
	}

	return txmodel.CreateWebhook(modeltarget)
}

func (s WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.modeladapt.AdaptTarget(target)

	txmodel, err := s.modelx.AutoTx()
	if err != nil {
		return result, err
	}

	modelresult, err := txmodel.GetWebhook(modeltarget)
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
