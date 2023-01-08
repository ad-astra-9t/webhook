package store

import (
	"context"
	"database/sql"

	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/model"
)

type Store interface {
	CreateWebhook(target domain.Webhook) error
	GetWebhook(target domain.Webhook) (result domain.Webhook, err error)
}

type DefaultStore struct {
	defaultModel *model.DefaultModel
	modelAdapt   *model.ModelAdapt
	WebhookStore
	EventStore
}

func (s *DefaultStore) Tx(ctx context.Context, options *sql.TxOptions) (*TxStore, error) {
	tx, err := s.defaultModel.Tx(ctx, options)
	if err != nil {
		return nil, err
	}

	return &TxStore{
		tx,
		s.modelAdapt,
		NewWebhookStore(tx, s.modelAdapt),
		NewEventStore(tx),
	}, nil
}

type TxStore struct {
	txModel    *model.TxModel
	modelAdapt *model.ModelAdapt
	WebhookStore
	EventStore
}

func (m *TxStore) Cancel() error {
	return m.txModel.Cancel()
}

func (m *TxStore) End() error {
	return m.txModel.End()
}

func NewDefaultStore(model *model.DefaultModel, adapt *model.ModelAdapt) *DefaultStore {
	return &DefaultStore{
		model,
		adapt,
		NewWebhookStore(model, adapt),
		NewEventStore(model),
	}
}
