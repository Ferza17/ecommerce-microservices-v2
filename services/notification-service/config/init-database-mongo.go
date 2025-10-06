package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type DatabaseMongo struct {
	Username                 string
	Password                 string
	Host                     string
	Port                     string
	DatabaseNameNotification string

	keyPrefix string
}

func DefaultDatabaseMongo() *DatabaseMongo {
	return &DatabaseMongo{
		keyPrefix: "%s/database/mongodb/%s",
	}
}

func (c *Config) withConfigDatabaseMongo(kv *api.KV) *Config {
	c.DatabaseMongo = DefaultDatabaseMongo().WithConsulClient(c.Env, kv)
	return c
}

func (c *DatabaseMongo) WithConsulClient(env string, kv *api.KV) *DatabaseMongo {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_USERNAME is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PORT is required")
	}
	c.Port = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_DATABASE_NAME/NOTIFICATION"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/NOTIFICATION from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/NOTIFICATION is required")
	}
	c.DatabaseNameNotification = string(pair.Value)

	return c
}
