package model

import "github.com/ad-astra-9t/webhook/dbx"

type Model struct {
	WebhookModel
	EventModel
}

func NewModel(dbx *dbx.DBX) *Model {
	return &Model{
		NewWebhookModel(dbx),
		NewEventModel(dbx),
	}
}
