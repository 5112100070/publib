package dummyrds_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/5112100070/publib/storage/redis"
	. "github.com/5112100070/publib/storage/redis/dummyrds"
)

func TestGet(t *testing.T) {

	m := Mocker{}
	m.AddMock("GET foo", "result", false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, "result", rds.Get("foo").String())
	assert.Equal(t, "", rds.Get("none").String())
}

func TestSetex(t *testing.T) {

	m := Mocker{}
	m.AddMock("SETEX foo 10 bar", "1", false)
	m.AddMock("SETEX err 10 bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Setex("foo", 10, "bar"))
	assert.EqualError(t, rds.Setex("err", 10, "bar"), "failed")
}

func TestDel(t *testing.T) {

	m := Mocker{}
	m.AddMock("DEL foo", "1", false)
	m.AddMock("DEL bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Del("foo"))
	assert.EqualError(t, rds.Del("bar"), "failed")
}
func TestExpire(t *testing.T) {

	m := Mocker{}
	m.AddMock("EXPIRE foo 10", "1", false)
	m.AddMock("EXPIRE err 10", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Expire("foo", 10))
	assert.EqualError(t, rds.Expire("err", 10), "failed")
}

func TestIncr(t *testing.T) {

	m := Mocker{}
	m.AddMock("INCR foo", "1", false)
	m.AddMock("INCR bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Incr("foo"))
	assert.EqualError(t, rds.Incr("bar"), "failed")
}

func TestIncrSingle(t *testing.T) {
	m := Mocker{}
	m.AddMock("INCR foo", "1", false)
	m.AddMock("INCR bar", "0", true)

	rds := New(Config{
		MockingMap: m,
	})

	res1, err1 := rds.IncrSingle("foo")
	res2, err2 := rds.IncrSingle("bar")

	assert.Equal(t, res1, 1)
	assert.Nil(t, err1)

	assert.Equal(t, res2, 0)
	assert.Nil(t, err2)
}

func TestDecr(t *testing.T) {

	m := Mocker{}
	m.AddMock("DECR foo", "1", false)
	m.AddMock("DECR bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Decr("foo"))
	assert.EqualError(t, rds.Decr("bar"), "failed")
}

func TestHDel(t *testing.T) {
	m := Mocker{}
	m.AddMock("HDEL test foo bar", "1", false)
	m.AddMock("HDEL test brown fox", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.HDel("test", "foo", "bar"))
	assert.EqualError(t, rds.HDel("test", "brown", "fox"), "failed")
}

func TestHDelSingle(t *testing.T) {
	m := Mocker{}
	m.AddMock("HDEL test foo", "1", false)
	m.AddMock("HDEL test bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.HDelSingle("test", "foo"))
	assert.EqualError(t, rds.HDelSingle("test", "bar"), "failed")
}

func TestHSet(t *testing.T) {

	m := Mocker{}
	m.AddMock("HSET foo bar res", "1", false)
	m.AddMock("HSET err bar res", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.HSet("foo", "bar", "res"))
	assert.EqualError(t, rds.HSet("err", "bar", "res"), "failed")
}

func TestHMSet(t *testing.T) {

	m := Mocker{}
	m.AddMock("HMSET foo bar res", "OK", false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.HMSet("foo", map[string]interface{}{"bar": "res"}))
}
func TestHGet(t *testing.T) {

	m := Mocker{}
	m.AddMock("HGET foo bar", "1", false)
	m.AddMock("HGET err bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 1, rds.HGet("foo", "bar").Int())
	assert.Equal(t, "", rds.HGet("err", "bar").String())
}

func TestHKeys(t *testing.T) {

	m := Mocker{}
	m.AddMock("HKEYS foo", []string{"one", "two"}, false)
	m.AddMock("HKEYS err", "", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Len(t, rds.HKeys("foo").StringSlice(), 2)
	assert.Equal(t, "", rds.HKeys("err").String())
}

func TestHVals(t *testing.T) {

	m := Mocker{}
	m.AddMock("HVALS foo", []string{"one", "two"}, false)
	m.AddMock("HVALS err", "", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Len(t, rds.HVals("foo").StringSlice(), 2)
	assert.Equal(t, "", rds.HVals("err").String())
}

func TestHGetALl(t *testing.T) {

	m := Mocker{}
	m.AddMock("HGETALL foo", []string{"one", "two"}, false)
	m.AddMock("HGETALL err", "", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Len(t, rds.HGetAll("foo").StringSlice(), 2)
	assert.Equal(t, "", rds.HGetAll("err").String())
}

func TestHExists(t *testing.T) {
	m := Mocker{}
	m.AddMock("HEXISTS key newKey", 0, false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 0, rds.HExists("key", "newKey").Int())
}

func Test_dummydis_ZRange(t *testing.T) {
	m := Mocker{}
	m.AddMock("ZRANGE foo", "0", false)
	m.AddMock("ZRANGE err", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 0, rds.ZRange("foo", 0, -1).Int())
	assert.Equal(t, "", rds.ZRange("err", 0, -1).String())
}

func Test_dummydis_ZRangeByScore(t *testing.T) {
	m := Mocker{}
	m.AddMock("ZRANGEBYSCORE foo", "0", false)
	m.AddMock("ZRANGEBYSCORE err", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 0, rds.ZRange("foo", 0, -1).Int())
	assert.Equal(t, "", rds.ZRange("err", 0, -1).String())
}

func Test_dummydis_Ttl(t *testing.T) {
	m := Mocker{}
	m.AddMock("TTL foo", "0", false)
	m.AddMock("TTL err", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 0, rds.Ttl("foo").Int())
	assert.Equal(t, "", rds.Ttl("err").String())
}

func Test_dummydis_ZAdd(t *testing.T) {
	m := Mocker{}
	m.AddMock("ZADD foo 1 test", "0", false)
	m.AddMock("ZADD err 1 test", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.ZAdd("foo", redis.Z{Score: 1, Member: "test"}))
	assert.EqualError(t, rds.ZAdd("err", redis.Z{Score: 1, Member: "test"}), "failed")
}

func Test_dummydis_Exists(t *testing.T) {
	m := Mocker{}
	m.AddMock("EXISTS aw", "0", false)
	m.AddMock("EXISTS error", "", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, rds.Exists("aw").String(), "0")
	assert.Equal(t, rds.Exists("error").String(), "")
}

func Test_dummydis_Rename(t *testing.T) {
	m := Mocker{}
	m.AddMock("RENAME key newKey", "OK", false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, rds.Rename("key", "newKey").String(), "OK")
}

func TestSet(t *testing.T) {

	m := Mocker{}
	m.AddMock("SET foo lol [NX]", "1", false)
	m.AddMock("SET err lol [EX]", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.Set("foo", "lol", "NX").Error)
	assert.EqualError(t, rds.Set("err", "lol", "EX").Error, "failed")
}

func TestLPush(t *testing.T) {

	m := Mocker{}
	m.AddMock("LPUSH foo bar", "1", false)
	m.AddMock("LPUSH err bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.LPush("foo", "bar"))
	assert.EqualError(t, rds.LPush("err", "bar"), "failed")
}

func TestRPush(t *testing.T) {

	m := Mocker{}
	m.AddMock("RPUSH foo bar", "1", false)
	m.AddMock("RPUSH err bar", "failed", true)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Nil(t, rds.RPush("foo", "bar"))
	assert.EqualError(t, rds.RPush("err", "bar"), "failed")
}

func TestLPop(t *testing.T) {

	m := Mocker{}
	m.AddMock("LPOP foo", "result", false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, "result", rds.LPop("foo").String())
	assert.Equal(t, "", rds.LPop("none").String())
}

func TestLLen(t *testing.T) {

	m := Mocker{}
	m.AddMock("LLEN mylist", 1, false)

	rds := New(Config{
		MockingMap: m,
	})

	assert.Equal(t, 1, rds.LLen("mylist").Int())
	assert.Equal(t, 0, rds.LLen("none").Int())
}
