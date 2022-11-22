package model

import "github.com/jmoiron/sqlx"

type DBX interface {
	ToDB() *sqlx.DB
	ToTx() *sqlx.Tx
	ToExt() (interface {
		sqlx.Ext
		sqlx.ExtContext
	}, error)
}
