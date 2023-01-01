package tx

import (
	"context"

	"github.com/ad-astra-9t/webhook/model"
	"github.com/ad-astra-9t/webhook/store"
)

type Storex struct {
	*ModelxStore
	Tx *ModelxStore
}

type ModelxStore struct {
	modelx *Modelx
	adapt  *model.ModelAdapt
	store.WebhookStore
	store.EventStore
}

func (s *Storex) SetTx(ctx context.Context) error {
	tx, err := s.ModelxStore.Tx(ctx)
	if err != nil {
		return err
	}

	s.Tx = tx

	return nil
}

func (s *Storex) Cancel() error {
	return s.modelx.Cancel()
}

func (s *Storex) End() error {
	return s.modelx.End()
}

func (s *ModelxStore) Tx(ctx context.Context) (*ModelxStore, error) {
	modelxCopy := new(Modelx)
	*modelxCopy = *s.modelx

	if err := modelxCopy.SetTx(ctx); err != nil {
		return nil, err
	}

	storeCopy := &ModelxStore{
		modelxCopy,
		s.adapt,
		store.NewWebhookStore(modelxCopy, s.adapt),
		store.NewEventStore(modelxCopy),
	}

	return storeCopy, nil
}

func NewModelxStore(modelx *Modelx, adapt *model.ModelAdapt) *ModelxStore {
	return &ModelxStore{
		modelx,
		adapt,
		store.NewWebhookStore(modelx, adapt),
		store.NewEventStore(modelx),
	}
}

func NewStorex(store *ModelxStore) *Storex {
	return &Storex{ModelxStore: store}
}
