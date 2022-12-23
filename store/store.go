package store

import (
	"github.com/ad-astra-9t/webhook/domain"
)

type Store interface {
	CreateWebhook(target domain.Webhook) error
	GetWebhook(target domain.Webhook) (result domain.Webhook, err error)
}
