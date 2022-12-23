package modelx

import "github.com/jmoiron/sqlx"

type Model interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

type DefaultModel struct {
	*sqlx.DB
	WebhookModel
	EventModel
}

func NewDefaultModel(db *sqlx.DB) *DefaultModel {
	return &DefaultModel{
		db,
		NewWebhookModel(db),
		NewEventModel(db),
	}
}
