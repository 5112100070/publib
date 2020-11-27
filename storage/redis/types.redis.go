package redis

// A Redis offers a standard interface for caching mechanism
type Redis interface {
	Ping() *Result
	Get(string) *Result
	MGet(...string) *Result
	Setex(string, int, interface{}) error
	Expire(string, int) error
	Del(...string) error
	HDel(string, ...string) error
	HDelSingle(string, string) error
	HSet(string, string, interface{}) error
	HMSet(key string, values map[string]interface{}) error
	HGet(string, string) *Result
	HKeys(hash string) *Result
	HVals(hash string) *Result
	HGetAll(hash string) *Result
	HExists(key string, fieldKey string) *Result
	Incr(...string) error
	IncrSingle(keys string) (int, error)
	Decr(...string) error
	ZAdd(string, ...Z) error
	ZRange(string, int, int) *Result
	ZRangeByScore(string, string, string, int) *Result
	Ttl(string) *Result
	Exists(string) *Result
	Rename(string, string) *Result
	Set(key, value interface{}, args ...interface{}) *Result
	LPush(key string, value interface{}) error
	RPush(key string, value interface{}) error
	LPop(key string) *Result
	LLen(key string) *Result
	Scan(cursor int, match string, count int) *Result
}

// Result struct
type Result struct {
	Value interface{}
	Error error
}
type Z struct {
	Score  float64
	Member interface{}
}
