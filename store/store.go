package store

import (
	"context"

	"github.com/ad-astra-9t/webhook/modelx"
)

type Store struct {
	modelx *modelx.Modelx
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
		NewWebhookStore(modelxCopy),
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

func NewStore(modelx *modelx.Modelx) *Store {
	return &Store{
		modelx,
		NewWebhookStore(modelx),
		NewEventStore(modelx),
	}
}
