package model

import (
	"github.com/ad-astra-9t/webhook/dbx"
)

type Model struct {
	dbx *dbx.DBX
	WebhookModel
	EventModel
}

func NewModel(dbx *dbx.DBX) *Model {
	return &Model{
		dbx,
		NewWebhookModel(dbx),
		NewEventModel(dbx),
	}
}
