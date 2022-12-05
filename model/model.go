package model

import (
	"context"

	"github.com/ad-astra-9t/webhook/dbx"
)

type Model struct {
	dbx *dbx.DBX
	WebhookModel
	EventModel
}

func (m *Model) Tx() *ModelTx {
	dbxCopy := new(dbx.DBX)
	*dbxCopy = *m.dbx
	modelCopy := &Model{
		dbxCopy,
		NewWebhookModel(dbxCopy),
		NewEventModel(dbxCopy),
	}
	return &ModelTx{modelCopy}
}

type ModelTx struct {
	*Model
}

func (c *ModelTx) Start(ctx context.Context) error {
	tx, err := c.dbx.BeginTxx(ctx, c.dbx.TxOptions)
	if err != nil {
		return err
	}

	c.dbx.Tx = tx

	return nil
}

func (c *ModelTx) Cancel() error {
	return c.dbx.Tx.Rollback()
}

func (c *ModelTx) End() error {
	return c.dbx.Tx.Commit()
}

func NewModel(dbx *dbx.DBX) *Model {
	return &Model{
		dbx,
		NewWebhookModel(dbx),
		NewEventModel(dbx),
	}
}
