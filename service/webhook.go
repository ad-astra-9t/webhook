package service

import "github.com/ad-astra-9t/webhook/domain"

type WebhookStore interface {
	CreateWebhook(domain.Webhook) error
	GetWebhook(domain.Webhook) (domain.Webhook, error)
}

type WebhookService struct {
	store WebhookStore
}

func (s *WebhookService) CreateWebhook(webhook domain.Webhook) error {
	return s.store.CreateWebhook(webhook)
}

func (s *WebhookService) GetWebhook(webhook domain.Webhook) (domain.Webhook, error) {
	return s.store.GetWebhook(webhook)
}
