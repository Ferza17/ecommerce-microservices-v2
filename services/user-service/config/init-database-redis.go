package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

type DatabaseRedis struct {
	Host      string
	Port      string
	Password  string
	DB        int
	keyPrefix string
}

func DefaultDatabaseRedis() *DatabaseRedis {
	return &DatabaseRedis{
		Host:      "",
		Port:      "",
		Password:  "",
		DB:        0,
		keyPrefix: "%s/database/redis/%s",
	}
}

func (c *Config) withDatabaseRedis(kv *api.KV) *Config {
	c.DatabaseRedis = DefaultDatabaseRedis().WithConsulClient(c.Env, kv)
	return c
}

func (c *DatabaseRedis) WithConsulClient(env string, kv *api.KV) *DatabaseRedis {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "REDIS_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "REDIS_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_PORT is required")
	}
	c.Port = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "REDIS_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get REDIS_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | REDIS_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "REDIS_DB"), nil)
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
	c.DB = int(redisDB)
	return c
}
