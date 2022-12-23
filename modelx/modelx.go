package modelx

import (
	"context"
)

type Modelx struct {
	*Model
	Tx *ModelTx
}

type TxModel interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

func (m *Modelx) SetTx(ctx context.Context) error {
	tx, err := m.Model.Tx(ctx)
	if err != nil {
		return err
	}

	m.Tx = tx

	return nil
}

func NewModelx(model *Model) *Modelx {
	return &Modelx{Model: model}
}
