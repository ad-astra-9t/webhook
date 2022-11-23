package store

import (
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/model"
)

type WebhookStore struct {
	model model.WebhookModel
}

func (s *WebhookStore) CreateWebhook(target domain.Webhook) error {
	modeltarget := s.model.AdaptModel(target)
	return s.model.CreateWebhook(modeltarget)
}

func (s *WebhookStore) GetWebhook(target domain.Webhook) (result domain.Webhook, err error) {
	modeltarget := s.model.AdaptModel(target)

	modelresult, err := s.model.GetWebhook(modeltarget)
	if err != nil {
		return result, err
	}

	result = s.model.AdaptDomain(modelresult)

	return result, err
}
