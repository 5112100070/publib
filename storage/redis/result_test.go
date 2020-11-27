package redis_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/5112100070/publib/storage/redis"
)

func TestInt(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Equal(t, 0, r.Int())

	r.Error = nil
	r.Value = 123
	assert.Equal(t, 123, r.Int())

	r.Value = int64(5000)
	assert.Equal(t, 5000, r.Int())

	r.Value = "foo bar"
	assert.Equal(t, 0, r.Int())

	r.Value = []byte("bar bar")
	assert.Equal(t, 0, r.Int())

}

func TestInt64(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Equal(t, int64(0), r.Int64())

	r.Error = nil
	r.Value = 123

	assert.Equal(t, int64(123), r.Int64())

	r.Value = int64(5000)
	assert.Equal(t, int64(5000), r.Int64())

	r.Value = "foo bar"
	assert.Equal(t, int64(0), r.Int64())

	r.Value = []byte("bar bar")
	assert.Equal(t, int64(0), r.Int64())

}

func TestString(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Equal(t, "", r.String())

	r.Error = nil
	r.Value = 123

	assert.Equal(t, "123", r.String())

	r.Value = int64(5000)
	assert.Equal(t, "5000", r.String())

	r.Value = "foo bar"
	assert.Equal(t, "foo bar", r.String())

	r.Value = []byte("bar bar")
	assert.Equal(t, "bar bar", r.String())

}

func TestStringSlice(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Equal(t, "", r.String())

	r.Error = nil
	r.Value = 123

	assert.Equal(t, []string(nil), r.StringSlice())

	r.Value = []string{"SFD"}
	assert.Equal(t, r.Value, r.StringSlice())

}

func TestBytes(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Nil(t, r.Bytes())

	r.Error = nil
	assert.Nil(t, r.Bytes())

	r.Value = 123
	assert.Equal(t, []byte("123"), r.Bytes())

	r.Value = int64(5000)
	assert.Equal(t, []byte("5000"), r.Bytes())

	r.Value = "foo bar"
	assert.Equal(t, []byte("foo bar"), r.Bytes())

	r.Value = []byte("bar bar")
	assert.Equal(t, []byte("bar bar"), r.Bytes())

}

func TestByteSlice(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	r.Error = nil
	assert.Nil(t, r.ByteSlice())

	r.Value = 123
	assert.Nil(t, r.ByteSlice())

	r.Value = int64(5000)
	assert.Nil(t, r.ByteSlice())

	r.Value = "foo bar"
	assert.Nil(t, r.ByteSlice())

	r.Value = []byte("bar bar")
	assert.Nil(t, r.ByteSlice())

}

func TestFloat64(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	assert.Equal(t, 0.0, r.Float64())

	r.Error = nil
	r.Value = 123
	assert.Equal(t, 123.0, r.Float64())

	r.Value = int64(5000)
	assert.Equal(t, 5000.0, r.Float64())

	r.Value = "foo bar"
	assert.Equal(t, 0.0, r.Float64())

	r.Value = []byte("bar bar")
	assert.Equal(t, 0.0, r.Float64())

}

func TestScanStruct(t *testing.T) {
	r := &Result{
		Error: errors.New("test"),
	}

	cursor, keys, err := r.ScanResult()
	assert.Equal(t, 0, cursor)
	assert.Equal(t, []string(nil), keys)
	assert.NotNil(t, err)
}
