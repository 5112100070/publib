package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Stmtx implement sqlx stmt
type Stmtx struct {
	db    *DB
	stmts []*sqlx.Stmt
}

// Close all dbs connection
func (st *Stmtx) Close() error {
	for i := range st.stmts {
		err := st.stmts[i].Close()

		if err != nil {
			return err
		}
	}
	return nil
}

// Exec will always go to production
func (st *Stmtx) Exec(args ...interface{}) (sql.Result, error) {
	return st.stmts[0].Exec(args...)

}

// Query will always go to slave
func (st *Stmtx) Query(args ...interface{}) (*sql.Rows, error) {
	return st.stmts[st.db.slave()].Query(args...)
}

// QueryMaster will use master db
func (st *Stmtx) QueryMaster(args ...interface{}) (*sql.Rows, error) {
	return st.stmts[0].Query(args...)
}

// QueryRow will always go to slave
func (st *Stmtx) QueryRow(args ...interface{}) *sql.Row {
	return st.stmts[st.db.slave()].QueryRow(args...)
}

// QueryRowMaster will use master db
func (st *Stmtx) QueryRowMaster(args ...interface{}) *sql.Row {
	return st.stmts[0].QueryRow(args...)
}

// MustExec using master database
func (st *Stmtx) MustExec(args ...interface{}) sql.Result {
	return st.stmts[0].MustExec(args...)
}

// Queryx will always go to slave
func (st *Stmtx) Queryx(args ...interface{}) (*sqlx.Rows, error) {
	return st.stmts[st.db.slave()].Queryx(args...)
}

// QueryRowx will always go to slave
func (st *Stmtx) QueryRowx(args ...interface{}) *sqlx.Row {
	return st.stmts[st.db.slave()].QueryRowx(args...)
}

// QueryRowxMaster will always go to master
func (st *Stmtx) QueryRowxMaster(args ...interface{}) *sqlx.Row {
	return st.stmts[0].QueryRowx(args...)
}

// Get will always go to slave
func (st *Stmtx) Get(dest interface{}, args ...interface{}) error {
	return st.stmts[st.db.slave()].Get(dest, args...)
}

// GetMaster will always go to master
func (st *Stmtx) GetMaster(dest interface{}, args ...interface{}) error {
	return st.stmts[0].Get(dest, args...)
}

// Select will always go to slave
func (st *Stmtx) Select(dest interface{}, args ...interface{}) error {
	return st.stmts[st.db.slave()].Select(dest, args...)
}

// SelectMaster will always go to master
func (st *Stmtx) SelectMaster(dest interface{}, args ...interface{}) error {
	return st.stmts[0].Select(dest, args...)
}

// ExecContext will always go to production
func (st *Stmtx) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return st.stmts[0].ExecContext(ctx, args...)
}

// QueryContext will always go to slave
func (st *Stmtx) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	return st.stmts[st.db.slave()].QueryContext(ctx, args...)
}

// QueryMasterContext will use master db
func (st *Stmtx) QueryMasterContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	return st.stmts[0].QueryContext(ctx, args...)
}

// QueryRowContext will always go to slave
func (st *Stmtx) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	return st.stmts[st.db.slave()].QueryRowContext(ctx, args...)
}

// QueryRowMasterContext will use master db
func (st *Stmtx) QueryRowMasterContext(ctx context.Context, args ...interface{}) *sql.Row {
	return st.stmts[0].QueryRowContext(ctx, args...)
}

// MustExecContext using master database
func (st *Stmtx) MustExecContext(ctx context.Context, args ...interface{}) sql.Result {
	return st.stmts[0].MustExecContext(ctx, args...)
}

// QueryxContext will always go to slave
func (st *Stmtx) QueryxContext(ctx context.Context, args ...interface{}) (*sqlx.Rows, error) {
	return st.stmts[st.db.slave()].QueryxContext(ctx, args...)
}

// QueryRowxContext will always go to slave
func (st *Stmtx) QueryRowxContext(ctx context.Context, args ...interface{}) *sqlx.Row {
	return st.stmts[st.db.slave()].QueryRowxContext(ctx, args...)
}

// QueryRowxMasterContext will always go to master
func (st *Stmtx) QueryRowxMasterContext(ctx context.Context, args ...interface{}) *sqlx.Row {
	return st.stmts[0].QueryRowxContext(ctx, args...)
}

// GetContext will always go to slave
func (st *Stmtx) GetContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return st.stmts[st.db.slave()].GetContext(ctx, dest, args...)
}

// GetMasterContext will always go to master
func (st *Stmtx) GetMasterContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return st.stmts[0].GetContext(ctx, dest, args...)
}

// SelectContext will always go to slave
func (st *Stmtx) SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return st.stmts[st.db.slave()].SelectContext(ctx, dest, args...)
}

// SelectMasterContext will always go to master
func (st *Stmtx) SelectMasterContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return st.stmts[0].SelectContext(ctx, dest, args...)
}
