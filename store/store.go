package store

import (
	"context"

	"github.com/ad-astra-9t/webhook/model"
	"github.com/ad-astra-9t/webhook/tx"
)

type Store struct {
	modelx *tx.Modelx
	adapt  *model.ModelAdapt
	WebhookStore
	EventStore
}

type StoreTx struct {
	*Store
}

func (s *Store) Tx(ctx context.Context) (*StoreTx, error) {
	modelxCopy := new(tx.Modelx)
	*modelxCopy = *s.modelx

	if err := modelxCopy.SetTx(ctx); err != nil {
		return nil, err
	}

	storeCopy := &Store{
		modelxCopy,
		s.adapt,
		NewWebhookStore(modelxCopy, s.adapt),
		NewEventStore(modelxCopy),
	}

	return &StoreTx{storeCopy}, nil
}

func (s *StoreTx) Cancel() error {
	return s.modelx.Tx.Cancel()
}

func (s *StoreTx) End() error {
	return s.modelx.Tx.End()
}

func NewStore(modelx *tx.Modelx, adapt *model.ModelAdapt) *Store {
	return &Store{
		modelx,
		adapt,
		NewWebhookStore(modelx, adapt),
		NewEventStore(modelx),
	}
}
