package model

import (
	"context"

	"github.com/ad-astra-9t/webhook/dbx"
)

type ModelChain struct {
	*Model
}

func (c *ModelChain) Start(ctx context.Context) error {
	tx, err := c.dbx.BeginTxx(ctx, c.dbx.TxOptions)
	if err != nil {
		return err
	}

	c.dbx.Tx = tx

	return nil
}

func (c *ModelChain) Cancel() error {
	return c.dbx.Tx.Rollback()
}

func (c *ModelChain) End() error {
	return c.dbx.Tx.Commit()
}

func NewModelChain(model *Model) ModelChain {
	dbxCopy := new(dbx.DBX)
	*dbxCopy = *model.dbx
	modelCopy := &Model{
		dbxCopy,
		NewWebhookModel(dbxCopy),
		NewEventModel(dbxCopy),
	}
	return ModelChain{modelCopy}
}