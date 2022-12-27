package tx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBX struct {
	*sqlx.DB
	Tx        *sqlx.Tx
	TxOptions *sql.TxOptions
}

func (d *DBX) SetTx(ctx context.Context) error {
	tx, err := d.BeginTxx(ctx, d.TxOptions)
	if err != nil {
		return err
	}

	d.Tx = tx

	return nil
}

func (d *DBX) Cancel() error {
	return d.Tx.Rollback()
}

func (d *DBX) End() error {
	return d.Tx.Commit()
}

func NewDBX(db *sqlx.DB) *DBX {
	return &DBX{DB: db}
}
