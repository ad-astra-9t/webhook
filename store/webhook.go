package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/modelx"
)

type WebhookStore struct {
	modelx *modelx.Modelx
}

func (s WebhookStore) CreateWebhook(target domain.Webhook) error {
	modeltarget := s.modelx.AdaptModel(target)

	txmodel, err := s.modelx.AutoTx()
	if err != nil {
		return err
	}

	return txmodel.CreateWebhook(modeltarget)
}

func (s WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.modelx.AdaptModel(target)

	txmodel, err := s.modelx.AutoTx()
	if err != nil {
		return result, err
	}

	modelresult, err := txmodel.GetWebhook(modeltarget)
	if err != nil {
		return result, err
	}

	result = s.modelx.AdaptDomain(modelresult)

	return result, err
}

func NewWebhookStore(modelx *modelx.Modelx) WebhookStore {
	return WebhookStore{modelx}
}
