package tx

import (
	"context"

	"github.com/ad-astra-9t/webhook/model"
)

type Modelx struct {
	*DBXModel
	Tx *ModelTx
}

type DBXModel struct {
	dbx *DBX
	model.WebhookModel
	model.EventModel
}

type ModelTx struct {
	*DBXModel
}

func (m *Modelx) SetTx(ctx context.Context) error {
	tx, err := m.DBXModel.Tx(ctx)
	if err != nil {
		return err
	}

	m.Tx = tx

	return nil
}

func (m *Modelx) Cancel() error {
	return m.Tx.Cancel()
}

func (m *Modelx) End() error {
	return m.Tx.End()
}

func (m *DBXModel) Tx(ctx context.Context) (*ModelTx, error) {
	dbxCopy := new(DBX)
	*dbxCopy = *m.dbx

	if err := dbxCopy.SetTx(ctx); err != nil {
		return nil, err
	}

	modelCopy := &DBXModel{
		dbxCopy,
		model.NewWebhookModel(dbxCopy),
		model.NewEventModel(dbxCopy),
	}

	return &ModelTx{modelCopy}, nil
}

func (m *ModelTx) Cancel() error {
	return m.dbx.Cancel()
}

func (m *ModelTx) End() error {
	return m.dbx.End()
}

func NewDBXModel(dbx *DBX) *DBXModel {
	return &DBXModel{
		dbx,
		model.NewWebhookModel(dbx),
		model.NewEventModel(dbx),
	}
}

func NewModelx(model *DBXModel) *Modelx {
	return &Modelx{DBXModel: model}
}
