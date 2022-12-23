package model

import "github.com/ad-astra-9t/webhook/domain"

type ModelAdapt struct {
	WebhookAdapt
}

type WebhookAdapt struct{}

func (a WebhookAdapt) AdaptTarget(domainwebhook domain.Webhook) (modelwebhook Webhook) {
	modelwebhook = Webhook{
		ID:       domainwebhook.ID,
		Callback: domainwebhook.Callback,
	}
	return
}

func (a WebhookAdapt) AdaptDomain(modelwebhook Webhook) (domainwebhook domain.Webhook) {
	domainwebhook = domain.Webhook{
		ID:       modelwebhook.ID,
		Callback: modelwebhook.Callback,
	}
	return
}
