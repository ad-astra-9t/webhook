package modelx

type Model interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}
