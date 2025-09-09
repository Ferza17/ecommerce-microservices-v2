package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func (c *Config) initRedis(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/database/redis/REDIS_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_HOST is required")
	}
	c.RedisHost = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_PORT is required")
	}
	c.RedisPort = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_PASSWORD", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_PASSWORD is required")
	}
	c.RedisPassword = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf("%s/database/redis/REDIS_DB", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_DB from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_DB is required")
	}
	redisDB, err := strconv.ParseInt(string(pair.Value), 10, 64)
	if err != nil {
		log.Fatalf("SetConfig | could not parse REDIS_DB to int: %v", err)
	}
	c.RedisDB = int(redisDB)
}
