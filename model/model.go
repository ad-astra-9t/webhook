package model

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Model interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

type DefaultModel struct {
	db *sqlx.DB
	WebhookModel
	EventModel
}

func (m *DefaultModel) Tx(ctx context.Context, options *sql.TxOptions) (*TxModel, error) {
	tx, err := m.db.BeginTxx(ctx, options)
	if err != nil {
		return nil, err
	}

	return &TxModel{
		tx,
		NewWebhookModel(tx),
		NewEventModel(tx),
	}, nil
}

type TxModel struct {
	tx *sqlx.Tx
	WebhookModel
	EventModel
}

func (m *TxModel) Cancel() error {
	return m.tx.Rollback()
}

func (m *TxModel) End() error {
	return m.tx.Commit()
}

func NewDefaultModel(db *sqlx.DB) *DefaultModel {
	return &DefaultModel{
		db,
		NewWebhookModel(db),
		NewEventModel(db),
	}
}
