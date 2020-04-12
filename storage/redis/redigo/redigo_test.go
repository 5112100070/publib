package redigo_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/5112100070/publib/storage/redis"
	. "github.com/5112100070/publib/storage/redis/redigo"
)

func TestGet(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	res := c.Get("test").String()
	assert.Equal(t, "", res)
}

func TestSetex(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Setex("test", 100, 1), "dial tcp: address null: missing port in address")
}
func TestDel(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Del("test"), "dial tcp: address null: missing port in address")
}

func TestExpire(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Expire("test", 123), "dial tcp: address null: missing port in address")
}

func TestIncr(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Incr("test"), "dial tcp: address null: missing port in address")
}

func TestIncrSingle(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	res, err := c.IncrSingle("test")

	assert.Equal(t, res, 0)
	assert.NotNil(t, err)
}

func TestDecr(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Decr("test"), "dial tcp: address null: missing port in address")
}

func TestHDel(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.HDel("test", "foo", "bar"), "dial tcp: address null: missing port in address")
}

func TestHDelSingle(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.HDelSingle("test", "bar"), "dial tcp: address null: missing port in address")
}

func TestHSet(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.HSet("test", "bar", 100), "dial tcp: address null: missing port in address")
}

func TestHMSet(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.HMSet("test", map[string]interface{}{"bar": 100}), "dial tcp: address null: missing port in address")
}

func TestHGet(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, "", c.HGet("test", "bar").String())
}

func TestHKeys(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, "", c.HKeys("test").String())
}

func TestHVals(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, "", c.HVals("test").String())
}

func TestHGetAll(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, []string(nil), c.HGetAll("test").StringSlice())
}

func TestHExists(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, 0, c.HExists("key", "newKey").Int())
}

func Test_credis_ZRange(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, []string(nil), c.ZRange("test", 0, -1).StringSlice())
}

func Test_credis_ZRangeByScore(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, []string(nil), c.ZRangeByScore("test", "-inf", "+inf", 0).StringSlice())
}

func Test_credis_Ttl(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.Equal(t, 0, c.Ttl("test").Int())
}

func Test_credis_ZAdd(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.ZAdd("test", redis.Z{Score: 1, Member: "test"}), "dial tcp: address null: missing port in address")
}

func Test_credis_Exists(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)
	assert.Equal(t, "", c.Exists("test").String())
}

func Test_credis_Rename(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)
	assert.Equal(t, "", c.Rename("key", "newKey").String())
}

func Test_credis_MGet(t *testing.T) {
	type fields struct {
		config Config
		// pool   rgo.Pool
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "happy case",
			fields: fields{
				config: Config{
					Endpoint: "null",
				},
			},
			args: args{
				key: "test",
			},
			want: []string(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.fields.config)
			got := c.MGet(tt.args.key)
			if !reflect.DeepEqual(got.StringSlice(), tt.want) {
				t.Errorf("credis.MGet() = %v, want %v", got.StringSlice(), tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.Set("test", 1, "EX", 180, "NX").Error, "dial tcp: address null: missing port in address")
}

func TestLPush(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.LPush("test", 1), "dial tcp: address null: missing port in address")
}

func TestRPush(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	assert.EqualError(t, c.RPush("test", 1), "dial tcp: address null: missing port in address")
}

func TestLPop(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	res := c.LPop("test").String()
	assert.Equal(t, "", res)
}

func TestLLen(t *testing.T) {

	cfg := Config{
		Endpoint: "null",
	}
	c := New(cfg)

	res := c.LLen("test").Int()
	assert.Equal(t, 0, res)
}

func TestScan(t *testing.T) {
	cfg := Config{
		Endpoint: "null",
	}

	c := New(cfg)

	assert.EqualError(t, c.Scan(0, "test", 8000).Error, "dial tcp: address null: missing port in address")
}
