package redigo

import (
	"fmt"
	"time"

	"github.com/5112100070/publib/convert"
	"github.com/5112100070/publib/storage/redis"
	rgo "github.com/gomodule/redigo/redis"
)

// New redis redigo module
func New(config Config) redis.Redis {

	// Set default 10 seconds timeout
	if config.Timeout == 0 {
		config.Timeout = 10
	}

	// Open connection to redis server
	return &credis{
		pool: rgo.Pool{
			MaxIdle:     config.MaxIdle,
			IdleTimeout: time.Duration(config.Timeout) * time.Second,
			Dial: func() (rgo.Conn, error) {
				return rgo.Dial(
					"tcp",
					config.Endpoint,
				)
			},
		},
	}

}

func (r *credis) Ping() *redis.Result {
	return r.cmd("PING")
}

func (c *credis) Get(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("GET", args...)
}

func (c *credis) MGet(keys ...string) *redis.Result {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return c.cmd("MGET", args...)
}

func (c *credis) Setex(key string, expireTime int, value interface{}) error {
	args := []interface{}{key, expireTime, value}
	return c.cmd("SETEX", args...).Error
}

func (c *credis) Del(keys ...string) error {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return c.cmd("DEL", args...).Error
}

func (c *credis) Expire(key string, seconds int) error {
	args := []interface{}{key, seconds}
	return c.cmd("EXPIRE", args...).Error
}

func (c *credis) Incr(keys ...string) error {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return c.cmd("INCR", args...).Error
}

func (c *credis) IncrSingle(keys string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("INCR", keys)
	r, err := conn.Do("EXEC")
	if err != nil {
		return 0, err
	}

	return convert.ToInt(r), nil
}

func (c *credis) Decr(keys ...string) error {
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	return c.cmd("DECR", args...).Error
}

// HDel is used to delete multiple fields
func (c *credis) HDel(key string, fields ...string) error {
	args := make([]interface{}, len(fields)+1)
	args[0] = key
	for i, v := range fields {
		args[i+1] = v
	}
	return c.cmd("HDEL", args...).Error
}

// HDel is used to delete single field
func (c *credis) HDelSingle(key, field string) error {
	args := []interface{}{key, field}
	return c.cmd("HDEL", args...).Error
}

func (c *credis) HSet(key, field string, value interface{}) error {
	args := []interface{}{key, field, value}
	return c.cmd("HSET", args...).Error
}

func (c *credis) HGet(key, field string) *redis.Result {
	args := []interface{}{key, field}
	return c.cmd("HGET", args...)
}

func (c *credis) HKeys(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("HKEYS", args...)
}

func (c *credis) HVals(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("HVALS", args...)
}

func (c *credis) HGetAll(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("HGETALL", args...)
}

func (c *credis) ZRange(key string, start int, end int) *redis.Result {
	args := []interface{}{key, start, end}
	return c.cmd("ZRANGE", args...)
}

func (c *credis) ZRangeByScore(key, min, max string, limit int) *redis.Result {
	args := []interface{}{key, min, max}
	if limit > 0 {
		args = []interface{}{key, min, max, "LIMIT", 0, limit}
	}
	return c.cmd("ZRANGEBYSCORE", args...)
}

func (c *credis) Ttl(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("TTL", args...)
}

func (c *credis) ZAdd(key string, values ...redis.Z) error {
	args := []interface{}{key}

	for _, value := range values {
		args = append(args, fmt.Sprintf("%v", value.Score), value.Member)
	}

	return c.cmd("ZADD", args...).Error
}

func (c *credis) Exists(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("EXISTS", args...)
}

func (c *credis) HMSet(key string, values map[string]interface{}) error {
	args := []interface{}{key}

	for key, val := range values {
		args = append(args, key, val)
	}

	result := c.cmd("HMSET", args...)
	return result.Error
}

func (c *credis) Rename(key, newKey string) *redis.Result {
	return c.cmd("RENAME", key, newKey)
}

func (c *credis) HExists(key string, fieldKey string) *redis.Result {
	return c.cmd("HEXISTS", key, fieldKey)
}

func (c *credis) cmd(command string, args ...interface{}) *redis.Result {
	result := &redis.Result{}
	conn := c.pool.Get()
	defer conn.Close()

	data, err := conn.Do(command, args...)
	if err != nil {

		// Retry mechanism
		conn2 := c.pool.Get()
		defer conn2.Close()
		data, err = conn2.Do(command, args...)
		if err != nil {
			result.Error = err
			return result
		}
	}
	result.Value = data

	return result
}

func (r *credis) Set(key, value interface{}, args ...interface{}) *redis.Result {
	args = append([]interface{}{key, value}, args...)
	return r.cmd("SET", args...)
}

// Insert all the specified values at the Head of the list stored at key
func (c *credis) LPush(key string, value interface{}) error {
	args := []interface{}{key, value}
	return c.cmd("LPUSH", args...).Error
}

// Insert all the specified values at the tail of the list stored at key
func (c *credis) RPush(key string, value interface{}) error {
	args := []interface{}{key, value}
	return c.cmd("RPUSH", args...).Error
}

// Removes and returns the first element of the list stored at key
func (c *credis) LPop(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("LPOP", args...)
}

// return length of element of the list
func (c *credis) LLen(key string) *redis.Result {
	args := []interface{}{key}
	return c.cmd("LLEN", args...)
}

// return cursor
func (c *credis) Scan(cursor int, match string, count int) *redis.Result {
	if count == 0 {
		count = 10
	}
	if match == "" {
		return c.cmd("SCAN", cursor, "COUNT", count)
	} else {
		return c.cmd("SCAN", cursor, "COUNT", count, "MATCH", match)
	}
}
