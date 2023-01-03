package cptx

import (
	"context"
	"errors"

	"github.com/ad-astra-9t/webhook/model"
)

type ModelCptx struct {
	model.WebhookModel
	model.EventModel
	dbCptx *DBCptx
}

func (c *ModelCptx) Cptx(ctx context.Context) (*ModelCptx, error) {
	dbCopy, err := c.dbCptx.Cptx(ctx)
	if err != nil {
		return nil, err
	}

	modelCopy := &ModelCptx{
		model.NewWebhookModel(dbCopy),
		model.NewEventModel(dbCopy),
		dbCopy,
	}

	return modelCopy, nil
}

func (c *ModelCptx) Cancel() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.dbCptx.Cancel()
}

func (c *ModelCptx) End() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.dbCptx.End()
}

func (c *ModelCptx) IsTx() bool {
	return c.dbCptx.IsTx()
}

func NewModelCptx(cptx *DBCptx) *ModelCptx {
	return &ModelCptx{
		model.NewWebhookModel(cptx),
		model.NewEventModel(cptx),
		cptx,
	}
}
