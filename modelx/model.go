package modelx

import (
	"context"

	"github.com/ad-astra-9t/webhook/tx"
)

type Model interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

type DBXModel struct {
	dbx *tx.DBX
	WebhookModel
	EventModel
}

type ModelTx struct {
	*DBXModel
}

func (m *DBXModel) Tx(ctx context.Context) (*ModelTx, error) {
	dbxCopy := new(tx.DBX)
	*dbxCopy = *m.dbx

	if err := dbxCopy.SetTx(ctx); err != nil {
		return nil, err
	}

	modelCopy := &DBXModel{
		dbxCopy,
		NewWebhookModel(dbxCopy),
		NewEventModel(dbxCopy),
	}

	return &ModelTx{modelCopy}, nil
}

func (m *ModelTx) Cancel() error {
	return m.dbx.Tx.Rollback()
}

func (m *ModelTx) End() error {
	return m.dbx.Tx.Commit()
}

func NewDBXModel(dbx *tx.DBX) *DBXModel {
	return &DBXModel{
		dbx,
		NewWebhookModel(dbx),
		NewEventModel(dbx),
	}
}
