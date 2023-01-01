package tx

import (
	"context"

	"github.com/ad-astra-9t/webhook/model"
)

type Modelx struct {
	*DBXModel
	swap *DBXModel
}

type DBXModel struct {
	dbx *DBX
	model.WebhookModel
	model.EventModel
}

func (m *Modelx) SetTx(ctx context.Context) error {
	tx, err := m.DBXModel.Tx(ctx)
	if err != nil {
		return err
	}

	m.swap = m.DBXModel
	m.DBXModel = tx

	return nil
}

func (m *Modelx) Cancel() error {
	tx := m.DBXModel
	m.DBXModel = m.swap
	m.swap = nil
	return tx.dbx.Cancel()
}

func (m *Modelx) End() error {
	tx := m.DBXModel
	m.DBXModel = m.swap
	m.swap = nil
	return tx.dbx.End()
}

func (m *DBXModel) Tx(ctx context.Context) (*DBXModel, error) {
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

	return modelCopy, nil
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
