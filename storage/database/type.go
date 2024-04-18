package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

// Error list
var (
	ErrNoConnectionDetected = errors.New("no connection detected")
)

// DB struct wrapper for sqlx connection
type DB struct {
	sqlxdb     []*sqlx.DB
	activedb   []int
	inactivedb []int
	dsn        []string
	driverName string
	groupName  string
	length     int
	count      uint64
	// for stats
	stats     []DbStatus
	heartBeat bool
	stopBeat  chan bool
	lastBeat  string
	// only use when needed
	debug bool
}

// DbStatus for status response
type DbStatus struct {
	Name       string      `json:"name"`
	Connected  bool        `json:"connected"`
	LastActive string      `json:"last_active"`
	Error      interface{} `json:"error"`
}

// Config database configuration
type Config struct {
	Driver string   `yaml:"driver"`
	Master string   `yaml:"master"`
	Slave  []string `yaml:"slave"`
}

type Database interface {
	// SetDebug for sql_wrapper
	SetDebug(v bool)
	// GetStatus return database status
	GetStatus() ([]DbStatus, error)
	// DoHeartBeat will automatically spawn a goroutines to ping your database every one second, use this carefully
	DoHeartBeat()
	// StopBeat will stop heartbeat, exit from goroutines
	StopBeat()
	// Ping database
	Ping() error
	// PingContext database
	PingContext(ctx context.Context) error
	// Prepare return sql stmt
	Prepare(query string) (Stmt, error)
	// Preparex sqlx stmt
	Preparex(query string) (*Stmtx, error)
	// PrepareContext return sql stmt
	PrepareContext(ctx context.Context, query string) (Stmt, error)
	// PreparexContext sqlx stmt
	PreparexContext(ctx context.Context, query string) (*Stmtx, error)
	// SetMaxOpenConnections to set max connections
	SetMaxOpenConnections(max int)
	// SetMaxIdleConnections to set max idle connections
	SetMaxIdleConnections(max int)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are reused forever.
	SetConnMaxLifetime(d time.Duration)
	// Slave return slave database
	Slave() *sqlx.DB
	// Master return master database
	Master() *sqlx.DB
	// Begin sql transaction
	Begin() (*sql.Tx, error)
	// Beginx sqlx transaction
	Beginx() (*sqlx.Tx, error)
	// MustBegin starts a transaction, and panics on error. Returns an *sqlx.Tx instead
	// of an *sql.Tx.
	MustBegin() *sqlx.Tx
	// BeginTx return sql.Tx
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	// Query queries the database and returns an *sql.Rows.
	// BeginTxx return sqlx.Tx
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryContext queries the database and returns an *sql.Rows.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow queries the database and returns an *sqlx.Row.
	QueryRow(query string, args ...interface{}) *sql.Row
	// QueryRowContext queries the database and returns an *sqlx.Row.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	// Queryx queries the database and returns an *sqlx.Rows.
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	// QueryxContext queries the database and returns an *sqlx.Rows.
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	// QueryRowx queries the database and returns an *sqlx.Row.
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	// QueryRowxContext queries the database and returns an *sqlx.Row.
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	// QueryRowxMasterContext queries the database and returns an *sqlx.Row.
	QueryRowxMasterContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	// Exec using master db
	Exec(query string, args ...interface{}) (sql.Result, error)
	// ExecContext using master db
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	// MustExec (panic) runs MustExec using master database.
	MustExec(query string, args ...interface{}) sql.Result
	// MustExecContext (panic) runs MustExec using master database.
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	// Select using slave db.
	Select(dest interface{}, query string, args ...interface{}) error
	// SelectContext using slave db.
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// SelectMaster using master db.
	SelectMaster(dest interface{}, query string, args ...interface{}) error
	// SelectMasterContext using master db.
	SelectMasterContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// Get using slave.
	Get(dest interface{}, query string, args ...interface{}) error
	// GetContext using slave.
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// GetMaster using master.
	GetMaster(dest interface{}, query string, args ...interface{}) error
	// GetMasterContext using master.
	GetMasterContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// NamedExec using master db.
	NamedExec(query string, arg interface{}) (sql.Result, error)
	// Rebind query
	Rebind(query string) string
	// RebindMaster will rebind query for master
	RebindMaster(query string) string
}
