package sql_wrapper

import (
	sql_wrapper_lib "github.com/5112100070/publib/database/sql_wrapper"
)

// Config database configuration
type Config struct {
	Driver string
	Master string
	Slave  []string
}

// database module
type sql_wrapperDB struct {
	config Config
	db     *sql_wrapper_lib.DB
}
