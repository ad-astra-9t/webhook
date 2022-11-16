package db

import "github.com/ad-astra-9t/webhook/domain"

type Webhook struct {
	ID       uint   `db:"id"`
	Callback string `db:"callback"`
}

func AdaptWebhook(webhook domain.Webhook) Webhook {
	return Webhook{
		ID:       webhook.ID,
		Callback: webhook.Callback,
	}
}
