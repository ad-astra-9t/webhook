package store

import (
	"context"

	"github.com/ad-astra-9t/webhook/modelx"
)

type Store struct {
	modelx *modelx.Modelx
	adapt  *modelx.ModelAdapt
	WebhookStore
	EventStore
}

type StoreTx struct {
	*Store
}

func (s *Store) Tx(ctx context.Context) (*StoreTx, error) {
	modelxCopy := new(modelx.Modelx)
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

func NewStore(modelx *modelx.Modelx, adapt *modelx.ModelAdapt) *Store {
	return &Store{
		modelx,
		adapt,
		NewWebhookStore(modelx, adapt),
		NewEventStore(modelx),
	}
}
