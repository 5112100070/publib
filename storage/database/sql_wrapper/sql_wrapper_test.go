package sql_wrapper_test

import (
	"testing"

	. "github.com/5112100070/publib/storage/database/sql_wrapper"
	"github.com/stretchr/testify/assert"
)

func TestNewFailed(t *testing.T) {

	db := New(Config{
		Driver: "mysql",
		Master: "master",
		Slave:  []string{"slave"},
	})

	assert.Nil(t, db)
}
