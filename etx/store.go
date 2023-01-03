package etx

import (
	"context"
	"errors"

	"github.com/ad-astra-9t/webhook/etx/cptx"
)

type StoreEtx struct {
	*cptx.StoreCptx
	exchange *cptx.StoreCptx
}

func (e *StoreEtx) Etx(ctx context.Context) error {
	if e.IsTx() {
		return errors.New("already in tx state")
	}

	tx, err := e.StoreCptx.Cptx(ctx)
	if err != nil {
		return err
	}

	e.exchange = e.StoreCptx
	e.StoreCptx = tx

	return nil
}

func (e *StoreEtx) Cancel() error {
	if !e.IsTx() {
		return errors.New("not in tx state")
	}

	tx := e.StoreCptx
	e.StoreCptx = e.exchange
	e.exchange = nil

	return tx.Cancel()
}

func (e *StoreEtx) End() error {
	if !e.IsTx() {
		return errors.New("not in tx state")
	}

	tx := e.StoreCptx
	e.StoreCptx = e.exchange
	e.exchange = nil

	return tx.End()
}

func NewStoreEtx(cptx *cptx.StoreCptx) *StoreEtx {
	return &StoreEtx{StoreCptx: cptx}
}
