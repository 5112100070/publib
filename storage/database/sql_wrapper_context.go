package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

// PingContext database
func (db *DB) PingContext(ctx context.Context) error {
	var err error

	if !db.heartBeat {
		for _, val := range db.sqlxdb {
			err = val.DB.PingContext(ctx)
			if err != nil {
				return err
			}
		}
		return err
	}

	for i := 0; i < len(db.activedb); i++ {
		val := db.activedb[i]
		err = db.sqlxdb[val].DB.PingContext(ctx)
		name := db.stats[val].Name

		if err != nil {
			if db.length <= 1 {
				return err
			}

			db.stats[val].Connected = false
			db.activedb = append(db.activedb[:i], db.activedb[i+1:]...)
			i--
			db.inactivedb = append(db.inactivedb, val)
			db.stats[val].Error = errors.New(name + ": " + err.Error())
			dbLengthMutex.Lock()
			db.length--
			dbLengthMutex.Unlock()
		} else {
			db.stats[val].Connected = true
			db.stats[val].LastActive = time.Now().Format(time.RFC1123)
			db.stats[val].Error = nil
		}
	}

	for i := 0; i < len(db.inactivedb); i++ {
		val := db.inactivedb[i]
		err = db.sqlxdb[val].DB.PingContext(ctx)
		name := db.stats[val].Name

		if err != nil {
			db.stats[val].Connected = false
			db.stats[val].Error = errors.New(name + ": " + err.Error())
		} else {
			db.stats[val].Connected = true
			db.inactivedb = append(db.inactivedb[:i], db.inactivedb[i+1:]...)
			i--
			db.activedb = append(db.activedb, val)
			db.stats[val].LastActive = time.Now().Format(time.RFC1123)
			db.stats[val].Error = nil
			dbLengthMutex.Lock()
			db.length++
			dbLengthMutex.Unlock()
		}
	}
	return err
}

// BeginTx return sql.Tx
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.Master().BeginTx(ctx, opts)
}

// BeginTxx return sqlx.Tx
func (db *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return db.Master().BeginTxx(ctx, opts)
}

// SelectContext using slave db.
func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.sqlxdb[db.slave()].SelectContext(ctx, dest, query, args...)
}

// SelectMasterContext using master db.
func (db *DB) SelectMasterContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.sqlxdb[0].SelectContext(ctx, dest, query, args...)
}

// GetContext using slave.
func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.sqlxdb[db.slave()].GetContext(ctx, dest, query, args...)
}

// GetMasterContext using master.
func (db *DB) GetMasterContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.sqlxdb[0].GetContext(ctx, dest, query, args...)
}

// PrepareContext return sql stmt
func (db *DB) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	var err error
	stmt := Stmt{}
	stmts := make([]*sql.Stmt, len(db.sqlxdb))

	for i := range db.sqlxdb {
		stmts[i], err = db.sqlxdb[i].PrepareContext(ctx, query)

		if err != nil {
			return stmt, err
		}
	}
	stmt.db = db
	stmt.stmts = stmts
	return stmt, nil
}

// PreparexContext sqlx stmt
func (db *DB) PreparexContext(ctx context.Context, query string) (*Stmtx, error) {
	var err error
	stmts := make([]*sqlx.Stmt, len(db.sqlxdb))

	for i := range db.sqlxdb {
		stmts[i], err = db.sqlxdb[i].PreparexContext(ctx, query)

		if err != nil {
			return nil, err
		}
	}

	return &Stmtx{db: db, stmts: stmts}, nil
}

// QueryContext queries the database and returns an *sql.Rows.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	r, err := db.sqlxdb[db.slave()].QueryContext(ctx, query, args...)
	return r, err
}

// QueryRowContext queries the database and returns an *sqlx.Row.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	rows := db.sqlxdb[db.slave()].QueryRowContext(ctx, query, args...)
	return rows
}

// QueryxContext queries the database and returns an *sqlx.Rows.
func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	r, err := db.sqlxdb[db.slave()].QueryxContext(ctx, query, args...)
	return r, err
}

// QueryRowxContext queries the database and returns an *sqlx.Row.
func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	rows := db.sqlxdb[db.slave()].QueryRowxContext(ctx, query, args...)
	return rows
}

// QueryRowxMasterContext queries the database and returns an *sqlx.Row.
func (db *DB) QueryRowxMasterContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	rows := db.sqlxdb[0].QueryRowxContext(ctx, query, args...)
	return rows
}

// ExecContext using master db
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.sqlxdb[0].ExecContext(ctx, query, args...)
}

// MustExecContext (panic) runs MustExec using master database.
func (db *DB) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return db.sqlxdb[0].MustExecContext(ctx, query, args...)
}
