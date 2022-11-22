package dbx

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type DBX struct {
	*sqlx.DB
	tx *sqlx.Tx
}

func (d *DBX) ToDB() *sqlx.DB {
	return d.DB
}

func (d *DBX) ToTx() *sqlx.Tx {
	return d.tx
}

func (d *DBX) ToExt() (interface {
	sqlx.Ext
	sqlx.ExtContext
}, error) {
	if d.tx != nil {
		return d.tx, nil
	}
	if d.DB != nil {
		return d.DB, nil
	}
	return nil, errors.New("failed to convert DBX to Ext")
}

func NewDBX(db *sqlx.DB, tx *sqlx.Tx) *DBX {
	return &DBX{DB: db, tx: tx}
}
