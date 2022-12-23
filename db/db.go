package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB interface {
	sqlx.Ext
	sqlx.ExtContext
}

func MustNewDB(driver string, dsn string) *sqlx.DB {
	return sqlx.MustOpen(driver, dsn)
}
