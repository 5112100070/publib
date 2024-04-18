package database

import "database/sql"

// Stmt implement sql stmt
type Stmt struct {
	db    *DB
	stmts []*sql.Stmt
}

// Exec will always go to production
func (st *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return st.stmts[0].Exec(args...)
}

// Query will always go to slave
func (st *Stmt) Query(args ...interface{}) (*sql.Rows, error) {
	return st.stmts[st.db.slave()].Query(args...)
}

// QueryMaster will use master db
func (st *Stmt) QueryMaster(args ...interface{}) (*sql.Rows, error) {
	return st.stmts[0].Query(args...)
}

// QueryRow will always go to slave
func (st *Stmt) QueryRow(args ...interface{}) *sql.Row {
	return st.stmts[st.db.slave()].QueryRow(args...)
}

// QueryRowMaster will use master db
func (st *Stmt) QueryRowMaster(args ...interface{}) *sql.Row {
	return st.stmts[0].QueryRow(args...)
}

// Close stmt
func (st *Stmt) Close() error {
	for i := range st.stmts {
		err := st.stmts[i].Close()

		if err != nil {
			return err
		}
	}
	return nil
}
