package model

import "errors"

type Modelx struct {
	*Model
	Tx *ModelTx
}

type AutoTxModel interface {
	GetWebhook(target Webhook) (result Webhook, err error)
	CreateWebhook(target Webhook) error
}

func (m *Modelx) AutoTxModel() (AutoTxModel, error) {
	if m.Tx != nil {
		return m.Tx, nil
	}
	if m.Model != nil {
		return m.Model, nil
	}
	return nil, errors.New("failed to convert Modelx to AutoTxModel")
}

func NewModelx(model *Model) *Modelx {
	return &Modelx{Model: model}
}
