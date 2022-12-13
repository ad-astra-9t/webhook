package modelx

import (
	"github.com/ad-astra-9t/webhook/dbx"
)

type Event struct {
	ID uint `db:"id"`
}

type EventModel struct {
	db dbx.DB
}

func NewEventModel(db dbx.DB) EventModel {
	return EventModel{db}
}
