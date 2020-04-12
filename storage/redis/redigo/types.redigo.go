package redigo

import (
	rgo "github.com/gomodule/redigo/redis"
)

type credis struct {
	// config Config
	pool   rgo.Pool
}

// Config of redis module
type Config struct {
	Endpoint string
	Timeout  int
	MaxIdle  int
}
