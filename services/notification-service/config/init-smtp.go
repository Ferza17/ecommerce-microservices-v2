package config

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type ConfigSmtp struct {
	SenderEmail string
	Host        string
	Port        string
	Username    string
	Password    string

	keyPrefix string
}

func DefaultConfigSmtp() *ConfigSmtp {
	return &ConfigSmtp{
		SenderEmail: "",
		Host:        "",
		Port:        "",
		Username:    "",
		Password:    "",
		keyPrefix:   "%s/smtp/%s",
	}
}

func (c *ConfigSmtp) WithConsulClient(env string, kv *api.KV) *ConfigSmtp {
	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "SMTP_SENDER_EMAIL"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_SENDER_EMAIL from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_SENDER_EMAIL is required")
	}
	c.SenderEmail = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SMTP_USERNAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_USERNAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_USERNAME is required")
	}
	c.Username = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SMTP_PASSWORD"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_PASSWORD from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_PASSWORD is required")
	}
	c.Password = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SMTP_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_HOST is required")
	}
	c.Host = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "SMTP_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_PORT is required")
	}
	c.Port = string(pair.Value)
	return c
}
