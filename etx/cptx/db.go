package cptx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/jmoiron/sqlx"
)

type DBCptx struct {
	db.DB
	db        *sqlx.DB
	txOptions *sql.TxOptions
}

func (c *DBCptx) Cptx(ctx context.Context) (*DBCptx, error) {
	tx, err := c.db.BeginTxx(ctx, c.txOptions)
	if err != nil {
		return nil, err
	}

	copy := &DBCptx{
		tx,
		c.db,
		c.txOptions,
	}

	return copy, nil
}

func (c *DBCptx) Cancel() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.DB.(*sqlx.Tx).Rollback()
}

func (c *DBCptx) End() error {
	if !c.IsTx() {
		return errors.New("not in tx state")
	}
	return c.DB.(*sqlx.Tx).Commit()
}

func (c *DBCptx) IsTx() bool {
	_, ok := c.DB.(*sqlx.Tx)
	return ok
}

func NewDBCptx(db *sqlx.DB, options *sql.TxOptions) *DBCptx {
	return &DBCptx{
		db,
		db,
		options,
	}
}
