package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type DatabaseMongodb struct {
	Username  string
	Password  string
	Host      string
	Port      string
	Database  string
	keyPrefix string
}

func DefaultDatabaseMongodb() *DatabaseMongodb {
	return &DatabaseMongodb{
		Username:  "",
		Password:  "",
		Host:      "",
		Port:      "",
		Database:  "",
		keyPrefix: "%s/database/mongodb/%s",
	}
}

func (c *DatabaseMongodb) WithConsulClient(env string, kv *api.KV) *DatabaseMongodb {

	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_USERNAME host is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PASSWORD host is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_HOST host is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/mongodb/MONGO_PORT", env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_PORT host is required")
	}
	c.Port = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "MONGO_DATABASE_NAME/EVENT_STORE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get MONGO_DATABASE_NAME/EVENT_STORE host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | MONGO_DATABASE_NAME/EVENT_STORE host is required")
	}
	c.Database = string(pair.Value)

	return c
}
