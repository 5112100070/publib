package dummydb

import (
	sql_wrapper_lib "github.com/5112100070/publib/database/sql_wrapper"
)

type dummyDB struct {
	db *sql_wrapper_lib.DB
}
