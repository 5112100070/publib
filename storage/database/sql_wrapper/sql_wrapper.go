package sql_wrapper

import (
	"database/sql"
	"log"
	"time"

	// MySQL
	_ "github.com/go-sql-driver/mysql"

	// PostgreSQL
	_ "github.com/lib/pq"

	sql_wrapper_lib "github.com/5112100070/publib/database/sql_wrapper"
	"github.com/5112100070/publib/storage/database"
	"github.com/jmoiron/sqlx"
)

// New sql_wrapper module
func New(config Config) database.Database {

	// Setup connection string
	connectionString := config.Master

	// Get slave value if available
	if config.Slave != nil {
		for _, v := range config.Slave {
			connectionString += ";" + v
		}
	}

	// Open connection to DB
	db, err := sql_wrapper_lib.Open(config.Driver, connectionString)
	if err != nil {
		log.Println("func New", err)
		return nil
	}

	db.SetMaxOpenConnections(100)
	db.SetConnMaxLifetime(time.Minute * 10)

	return &sql_wrapperDB{
		config: config,
		db:     db,
	}
}

func (f *sql_wrapperDB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return f.db.Queryx(query, args...)
}

func (f *sql_wrapperDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return f.db.Exec(query, args...)
}
func (f *sql_wrapperDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return f.db.Query(query, args...)
}

func (f *sql_wrapperDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return f.db.QueryRow(query, args...)
}

func (f *sql_wrapperDB) Begin() (*sql.Tx, error) {
	return f.db.Begin()
}

func (f *sql_wrapperDB) Beginx() (*sqlx.Tx, error) {
	return f.db.Beginx()
}
func (f *sql_wrapperDB) Master() *sqlx.DB {
	return f.db.Master()
}

func (f *sql_wrapperDB) Get(dest interface{}, query string, args ...interface{}) error {
	return f.db.Get(dest, query, args...)
}
