package autotx

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type DBX struct {
	*sqlx.DB
	Tx        *sqlx.Tx
	TxOptions *sql.TxOptions
}

type TxDB interface {
	sqlx.Ext
	sqlx.ExtContext
}

func (d *DBX) SetTx(ctx context.Context) error {
	tx, err := d.BeginTxx(ctx, d.TxOptions)
	if err != nil {
		return err
	}

	d.Tx = tx

	return nil
}

func (d *DBX) AutoTx() (TxDB, error) {
	if d.Tx != nil {
		return d.Tx, nil
	}
	if d.DB != nil {
		return d.DB, nil
	}
	return nil, errors.New("failed to convert DBX to TxDB")
}

func NewDBX(db *sqlx.DB) *DBX {
	return &DBX{DB: db}
}
