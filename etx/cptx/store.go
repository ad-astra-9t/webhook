package cptx

import (
	"context"
	"errors"

	"github.com/ad-astra-9t/webhook/model"
	"github.com/ad-astra-9t/webhook/store"
)

type StoreCptx struct {
	store.WebhookStore
	store.EventStore
	modelCptx *ModelCptx
	adapt     *model.ModelAdapt
}

func (c *StoreCptx) Cptx(ctx context.Context) (*StoreCptx, error) {
	modelCopy, err := c.modelCptx.Cptx(ctx)
	if err != nil {
		return nil, err
	}

	storeCopy := &StoreCptx{
		store.NewWebhookStore(modelCopy, c.adapt),
		store.NewEventStore(modelCopy),
		modelCopy,
		c.adapt,
	}

	return storeCopy, nil
}

func (c *StoreCptx) Cancel() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.modelCptx.Cancel()
}

func (c *StoreCptx) End() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.modelCptx.End()
}

func (c *StoreCptx) IsTx() bool {
	return c.modelCptx.IsTx()
}

func NewStoreCptx(cptx *ModelCptx, adapt *model.ModelAdapt) *StoreCptx {
	return &StoreCptx{
		store.NewWebhookStore(cptx, adapt),
		store.NewEventStore(cptx),
		cptx,
		adapt,
	}
}
