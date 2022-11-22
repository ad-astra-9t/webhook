package dbx

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustNewDB(driver string, dsn string) *sqlx.DB {
	return sqlx.MustOpen(driver, dsn)
}
