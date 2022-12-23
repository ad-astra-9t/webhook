package modelx

import (
	"github.com/ad-astra-9t/webhook/db"
)

type Event struct {
	ID uint `db:"id"`
}

type EventModel struct {
	db db.DB
}

func NewEventModel(db db.DB) EventModel {
	return EventModel{db}
}
