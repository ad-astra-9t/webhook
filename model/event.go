package model

type Event struct {
	ID uint `db:"id"`
}

type EventModel struct {
	autoTxDB AutoTxDB
}

func NewEventModel(autoTxDB AutoTxDB) EventModel {
	return EventModel{autoTxDB}
}
