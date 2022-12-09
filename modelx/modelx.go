package modelx

import "errors"

type Modelx struct {
	*Model
	Tx *ModelTx
}

type TxModel interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

func (m *Modelx) AutoTx() (TxModel, error) {
	if m.Tx != nil {
		return m.Tx, nil
	}
	if m.Model != nil {
		return m.Model, nil
	}
	return nil, errors.New("failed to convert Modelx to TxModel")
}

func NewModelx(model *Model) *Modelx {
	return &Modelx{Model: model}
}
