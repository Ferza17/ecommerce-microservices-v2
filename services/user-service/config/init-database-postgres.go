package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type DatabasePostgres struct {
	Username  string
	Password  string
	SSLMode   string
	Host      string
	Port      string
	Database  string
	keyPrefix string
}

func DefaultDatabasePostgres() *DatabasePostgres {
	return &DatabasePostgres{
		Username:  "",
		Password:  "",
		SSLMode:   "",
		Host:      "",
		Port:      "",
		Database:  "",
		keyPrefix: "%s/database/postgres/%s",
	}
}

func (c *Config) withDatabasePostgres(kv *api.KV) *Config {
	c.DatabasePostgres = DefaultDatabasePostgres().WithConsulClient(c.Env, kv)
	return c
}

func (c *DatabasePostgres) WithConsulClient(env string, kv *api.KV) *DatabasePostgres {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_USERNAME is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_SSL_MODE"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_SSL_MODE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_SSL_MODE is required")
	}
	c.SSLMode = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_PORT is required")
	}
	c.Port = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "POSTGRES_DATABASE_NAME/USERS"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_DATABASE_NAME/USERS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_DATABASE_NAME/USERS is required")
	}
	c.Database = string(pair.Value)
	return c
}
