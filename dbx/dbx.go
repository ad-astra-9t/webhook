package dbx

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type DBX struct {
	*sqlx.DB
	Tx        *sqlx.Tx
	TxOptions *sql.TxOptions
}

type AutoTxDB interface {
	sqlx.Ext
	sqlx.ExtContext
}

func (d *DBX) AutoTxDB() (AutoTxDB, error) {
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
