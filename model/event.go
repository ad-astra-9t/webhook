package model

import "github.com/ad-astra-9t/webhook/dbx"

type Event struct {
	ID uint `db:"id"`
}

type EventModel struct {
	dbx *dbx.DBX
}

func NewEventModel(dbx *dbx.DBX) EventModel {
	return EventModel{
		dbx: dbx,
	}
}
