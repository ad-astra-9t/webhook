package modelx

import (
	"context"
)

type Modelx struct {
	*DBXModel
	Tx *ModelTx
}

func (m *Modelx) SetTx(ctx context.Context) error {
	tx, err := m.DBXModel.Tx(ctx)
	if err != nil {
		return err
	}

	m.Tx = tx

	return nil
}

func NewModelx(model *DBXModel) *Modelx {
	return &Modelx{DBXModel: model}
}
