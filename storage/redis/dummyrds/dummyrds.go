package dummyrds

import (
	"errors"
	"fmt"
	"strings"

	"github.com/5112100070/publib/storage/redis"
)

// New module for mocking
func New(config Config) redis.Redis {
	return &dummydis{
		config: config,
	}
}

// AddMock to mocking mapping
func (m Mocker) AddMock(command string, result interface{}, isError bool) {
	m[command] = mock{
		IsError: isError,
		Result:  result,
	}
}

func (c *dummydis) Get(key string) *redis.Result {
	return c.mock("GET " + key)
}

func (c *dummydis) MGet(keys ...string) *redis.Result {
	f := strings.Join(keys, " ")
	return c.mock(fmt.Sprintf("MGET %s", f))
}

func (c *dummydis) Setex(key string, expireTime int, value interface{}) error {
	return c.mock(fmt.Sprintf("SETEX %s %d %s", key, expireTime, value)).Error
}

func (c *dummydis) Del(keys ...string) error {
	k := strings.Join(keys, " ")
	return c.mock("DEL " + k).Error
}

func (c *dummydis) Expire(key string, seconds int) error {
	return c.mock(fmt.Sprintf("EXPIRE %s %d", key, seconds)).Error
}

func (c *dummydis) Incr(keys ...string) error {
	k := strings.Join(keys, " ")
	return c.mock("INCR " + k).Error
}

func (c *dummydis) IncrSingle(key string) (int, error) {
	return c.mock("INCR " + key).Int(), nil
}

func (c *dummydis) Decr(keys ...string) error {
	k := strings.Join(keys, " ")
	return c.mock("DECR " + k).Error
}

func (c *dummydis) HDel(key string, fields ...string) error {
	f := strings.Join(fields, " ")
	return c.mock(fmt.Sprintf("HDEL %s %s", key, f)).Error
}

func (c *dummydis) HDelSingle(key, field string) error {
	return c.mock(fmt.Sprintf("HDEL %s %s", key, field)).Error
}

func (c *dummydis) HSet(key, field string, value interface{}) error {
	return c.mock(fmt.Sprintf("HSET %s %s %s", key, field, value)).Error
}

func (c *dummydis) HMSet(key string, values map[string]interface{}) error {
	return nil
}

func (c *dummydis) HGet(key, field string) *redis.Result {
	return c.mock(fmt.Sprintf("HGET %s %s", key, field))
}

func (c *dummydis) HKeys(hash string) *redis.Result {
	return c.mock(fmt.Sprintf("HKEYS %s", hash))
}

func (c *dummydis) HVals(hash string) *redis.Result {
	return c.mock(fmt.Sprintf("HVALS %s", hash))
}

func (c *dummydis) HGetAll(hash string) *redis.Result {
	return c.mock(fmt.Sprintf("HGETALL %s", hash))
}

func (c *dummydis) HExists(key, newKey string) *redis.Result {
	return c.mock(fmt.Sprintf("HEXISTS %s %s", key, newKey))
}

func (c *dummydis) ZRange(key string, start int, end int) *redis.Result {
	return c.mock(fmt.Sprintf("ZRANGE %s %v %v", key, start, end))
}

func (c *dummydis) ZRangeByScore(key, min, max string, limit int) *redis.Result {
	if limit > 0 {
		return c.mock(fmt.Sprintf("ZRANGEBYSCORE %s %v %v LIMIT 0 %d", key, min, max, limit))
	} else {
		return c.mock(fmt.Sprintf("ZRANGEBYSCORE %s %v %v", key, min, max))
	}
}

func (c *dummydis) Ttl(hash string) *redis.Result {
	return c.mock(fmt.Sprintf("TTL %s", hash))
}

func (c *dummydis) ZAdd(key string, values ...redis.Z) error {
	req := fmt.Sprintf("ZADD %s", key)
	for _, value := range values {
		req = fmt.Sprintf("%s %v %v", req, value.Score, value.Member)
	}

	return c.mock(req).Error
}

func (c *dummydis) Exists(hash string) *redis.Result {
	return c.mock(fmt.Sprintf("EXISTS %s", hash))
}

func (c *dummydis) Rename(key, newKey string) *redis.Result {
	return c.mock(fmt.Sprintf("RENAME %s %s", key, newKey))
}

func (c *dummydis) mock(command string) *redis.Result {

	res, ok := c.config.MockingMap[command]

	// Mock not found
	if !ok {
		return &redis.Result{
			Error: errors.New("No mocking found for " + command),
		}
	}

	// Mock error
	if res.IsError {
		return &redis.Result{
			Error: fmt.Errorf("%v", res.Result),
		}
	}

	// Mock result
	return &redis.Result{
		Value: res.Result,
	}

}

func (c *dummydis) Set(key, value interface{}, args ...interface{}) *redis.Result {
	return c.mock(fmt.Sprintf("SET %s %s %s", key, value, args))
}

func (c *dummydis) LPush(key string, value interface{}) error {
	return c.mock(fmt.Sprintf("LPUSH %s %s", key, value)).Error
}

func (c *dummydis) RPush(key string, value interface{}) error {
	return c.mock(fmt.Sprintf("RPUSH %s %s", key, value)).Error
}

func (c *dummydis) LPop(key string) *redis.Result {
	return c.mock("LPOP " + key)
}

func (c *dummydis) LLen(key string) *redis.Result {
	return c.mock("LLEN " + key)
}

func (c *dummydis) Scan(cursor int, match string, count int) *redis.Result {
	return c.mock(fmt.Sprintf("SCAN %d %s %d", cursor, match, count))
}
