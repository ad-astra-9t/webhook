package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/model"
)

type WebhookStore struct {
	model        model.WebhookModel
	modelAdapter model.WebhookAdapter
}

func (s *WebhookStore) CreateWebhook(target domain.Webhook) error {
	modeltarget := s.modelAdapter.AdaptDB(target)
	return s.model.CreateWebhook(modeltarget)
}

func (s *WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.modelAdapter.AdaptDB(target)

	modelresult, err := s.model.GetWebhook(modeltarget)
	if err != nil {
		return result, err
	}

	result = s.modelAdapter.AdaptDomain(modelresult)

	return result, err
}
