package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initDatabasePostgres(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_USERNAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_USERNAME is required")
	}
	c.PostgresUsername = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PASSWORD", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_PASSWORD is required")
	}
	c.PostgresPassword = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_SSL_MODE", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_SSL_MODE from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_SSL_MODE is required")
	}
	c.PostgresSSLMode = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_HOST is required")
	}
	c.PostgresHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_PORT is required")
	}
	c.PostgresPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get POSTGRES_DATABASE_NAME/PAYMENTS from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | POSTGRES_DATABASE_NAME/PAYMENTS is required")
	}
	c.PostgresDatabaseName = string(pair.Value)
}
