package database

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Init database connection
// Call this function first before using database module
func Init(config map[string]*Config) map[string]Database {
	var dbConn = make(map[string]Database)

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
		db, err := Open(config[key].Driver, connectionString)
		if err != nil {
			continue
		}
		db.SetMaxOpenConnections(100)
		db.SetConnMaxLifetime(time.Minute * 10)

		dbConn[key] = db
	}

	return dbConn

}

func openContextConnection(ctx context.Context, driverName, sources string, groupName string) (*DB, error) {
	var err error

	conns := strings.Split(sources, ";")
	connsLength := len(conns)

	// check if no source is available
	if connsLength < 1 {
		return nil, errors.New("no sources found")
	}

	db := &DB{
		sqlxdb: make([]*sqlx.DB, connsLength),
		stats:  make([]DbStatus, connsLength),
	}
	db.length = connsLength
	db.driverName = driverName

	for i := range conns {
		db.sqlxdb[i], err = sqlx.Open(driverName, conns[i])
		if err != nil {
			db.inactivedb = append(db.inactivedb, i)
			return nil, err
		}
		constatus := true

		// set the name
		name := ""
		if i == 0 {
			name = "master"
		} else {
			name = "slave-" + strconv.Itoa(i)
		}

		status := DbStatus{
			Name:       name,
			Connected:  constatus,
			LastActive: time.Now().String(),
		}

		db.stats[i] = status
		db.activedb = append(db.activedb, i)
	}

	// set the default group name
	db.groupName = defaultGroupName
	if groupName != "" {
		db.groupName = groupName
	}

	// ping database to retrieve error
	err = db.PingContext(ctx)
	return db, err
}

// OpenWithContext opening connection with context
func OpenWithContext(ctx context.Context, driver, sources string) (*DB, error) {
	return openContextConnection(ctx, driver, sources, "")
}

const defaultGroupName = "sql_wrapper_open"

var dbLengthMutex = &sync.Mutex{}

func openConnection(driverName, sources string, groupName string) (*DB, error) {
	var err error

	conns := strings.Split(sources, ";")
	connsLength := len(conns)

	// check if no source is available
	if connsLength < 1 {
		return nil, errors.New("no sources found")
	}

	db := &DB{
		sqlxdb: make([]*sqlx.DB, connsLength),
		stats:  make([]DbStatus, connsLength),
		dsn:    make([]string, connsLength),
	}
	db.length = connsLength
	db.driverName = driverName

	for i := range conns {
		db.sqlxdb[i], err = sqlx.Open(driverName, conns[i])
		if err != nil {
			db.inactivedb = append(db.inactivedb, i)
			return nil, err
		}

		constatus := true

		// set the name
		name := ""
		if i == 0 {
			name = "master"
		} else {
			name = "slave-" + strconv.Itoa(i)
		}

		status := DbStatus{
			Name:       name,
			Connected:  constatus,
			LastActive: time.Now().String(),
		}

		db.stats[i] = status
		db.activedb = append(db.activedb, i)
		db.dsn[i] = conns[i]
	}

	// set the default group name
	db.groupName = defaultGroupName
	if groupName != "" {
		db.groupName = groupName
	}

	// ping database to retrieve error
	err = db.Ping()
	return db, err
}

// Open connection to database
func Open(driverName, sources string) (*DB, error) {
	return openConnection(driverName, sources, "")
}

// OpenWithName open the connection and set connection group name
func OpenWithName(driverName, sources string, name string) (*DB, error) {
	return openConnection(driverName, sources, name)
}
