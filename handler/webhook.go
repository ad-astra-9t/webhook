package handler

import (
	"github.com/ad-astra-9t/webhook/domain"
)

type WebhookService interface {
	CreateWebhook(domain.Webhook) error
	GetWebhook(domain.Webhook) (domain.Webhook, error)
}

type WebhookHandler struct {
	service WebhookService
}
