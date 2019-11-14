package database

import (
	"log"
	"time"

	"github.com/5112100070/publib/database/sql_wrapper"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Config database configuration
type Config struct {
	Driver string
	Master string
	Slave  []string
}

var (
	sql_wrapperOpen = sql_wrapper.Open
)

// Init database connection
// Call this function first before using database module
func Init(config map[string]*Config) map[string]*sql_wrapper.DB {

	var dbConn = make(map[string]*sql_wrapper.DB)

	// Loop all given configurations
	for key := range config {

		connectionString := config[key].Master

		// Get slave value if available
		if config[key].Slave != nil {
			for _, v := range config[key].Slave {
				connectionString += ";" + v
			}
		}

		// Open connection to DB
		db, err := sql_wrapperOpen(config[key].Driver, connectionString)
		if err != nil {
			log.Println("func Init", err)
			continue
		}
		db.SetMaxOpenConnections(100)
		db.SetConnMaxLifetime(time.Minute * 10)

		dbConn[key] = db
	}

	return dbConn

}

// InitMock database
func InitMock() (*sql_wrapper.DB, sqlmock.Sqlmock) {
	db, mocker, _ := sqlmock.New()
	return sql_wrapper.InitMocking(db, 1), mocker
}
