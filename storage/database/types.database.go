package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// A Database interface provides connectivity to DB
type Database interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Begin() (*sql.Tx, error)
	Beginx() (*sqlx.Tx, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Master() *sqlx.DB
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
