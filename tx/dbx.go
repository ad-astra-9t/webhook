package tx

import (
	"context"
	"database/sql"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/jmoiron/sqlx"
)

type DBX struct {
	*TxDB
	Tx *TxDB
}

type TxDB struct {
	db        *sqlx.DB
	txOptions *sql.TxOptions
	db.DB
}

func (d *DBX) SetTx(ctx context.Context) error {
	tx, err := d.TxDB.Tx(ctx)
	if err != nil {
		return err
	}

	d.Tx = d.TxDB
	d.TxDB = tx

	return nil
}

func (d *DBX) Cancel() error {
	return d.TxDB.DB.(*sqlx.Tx).Rollback()
}

func (d *DBX) End() error {
	return d.TxDB.DB.(*sqlx.Tx).Commit()
}

func (t *TxDB) Tx(ctx context.Context) (*TxDB, error) {
	tx, err := t.db.BeginTxx(ctx, t.txOptions)
	if err != nil {
		return nil, err
	}

	dbCopy := &TxDB{
		t.db,
		t.txOptions,
		tx,
	}

	return dbCopy, nil
}

func NewTxDB(db *sqlx.DB, options *sql.TxOptions) *TxDB {
	return &TxDB{
		db,
		options,
		db,
	}
}

func NewDBX(db *TxDB) *DBX {
	return &DBX{TxDB: db}
}
