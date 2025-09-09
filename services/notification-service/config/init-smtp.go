package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func (c *Config) initSmtp(kv *api.KV) {
	pair, _, err := kv.Get(fmt.Sprintf("%s/smtp/SMTP_SENDER_EMAIL", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_SENDER_EMAIL host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_SENDER_EMAIL host is required")
	}
	c.SmtpSenderEmail = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_USERNAME", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_USERNAME host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_USERNAME host is required")
	}
	c.SmtpUsername = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_PASSWORD", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_PASSWORD host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_PASSWORD host is required")
	}
	c.SmtpPassword = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_HOST", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_HOST host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_HOST host is required")
	}
	c.SmtpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf("%s/smtp/SMTP_PORT", c.Env), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SMTP_PORT host from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SMTP_PORT host is required")
	}
	c.SmtpPort = string(pair.Value)
}
