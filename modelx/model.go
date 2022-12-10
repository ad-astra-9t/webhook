package modelx

import (
	"context"

	"github.com/ad-astra-9t/webhook/dbx"
)

type Model struct {
	dbx *dbx.DBX
	WebhookModel
	EventModel
}

type ModelTx struct {
	*Model
}

type AutoTxDB interface {
	AutoTx() (dbx.TxDB, error)
}

func (m *Model) Tx(ctx context.Context) (*ModelTx, error) {
	dbxCopy := new(dbx.DBX)
	*dbxCopy = *m.dbx

	if err := dbxCopy.SetTx(ctx); err != nil {
		return nil, err
	}

	modelCopy := &Model{
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

func NewModel(dbx *dbx.DBX) *Model {
	return &Model{
		dbx,
		NewWebhookModel(dbx),
		NewEventModel(dbx),
	}
}
