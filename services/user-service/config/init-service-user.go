package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/xhit/go-str2duration/v2"
	"log"
	"time"
)

type ConfigServiceUser struct {
	ServiceName string

	RpcHost        string
	RpcPort        string
	HttpHost       string
	HttpPort       string
	MetricHttpPort string

	JwtAccessTokenSecret          string
	JwtAccessTokenExpirationTime  time.Duration
	JwtRefreshTokenSecret         string
	JwtRefreshTokenExpirationTime time.Duration
	VerificationUserLoginUrl      string
	OtpExpirationTime             time.Duration

	keyPrefix string
}

func DefaultConfigServiceUser() *ConfigServiceUser {
	return &ConfigServiceUser{
		ServiceName:                   "",
		RpcHost:                       "",
		RpcPort:                       "",
		HttpHost:                      "",
		HttpPort:                      "",
		MetricHttpPort:                "",
		JwtAccessTokenSecret:          "",
		JwtAccessTokenExpirationTime:  0,
		JwtRefreshTokenSecret:         "",
		JwtRefreshTokenExpirationTime: 0,
		VerificationUserLoginUrl:      "",
		OtpExpirationTime:             0,
		keyPrefix:                     "%s/services/user/%s",
	}
}

func (c *Config) withServiceUser(kv *api.KV) *Config {
	c.ConfigServiceUser = DefaultConfigServiceUser().WithConsulClient(c.Env, kv)
	return c
}

func (c *ConfigServiceUser) WithConsulClient(env string, kv *api.KV) *ConfigServiceUser {

	pair, _, err := kv.Get(fmt.Sprintf(c.keyPrefix, env, "SERVICE_NAME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get SERVICE_NAME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | SERVICE_NAME is required")
	}
	c.ServiceName = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RPC_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_HOST is required")
	}
	c.RpcHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "RPC_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get RPC_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | RPC_PORT is required")
	}
	c.RpcPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "HTTP_HOST"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_HOST from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_HOST is required")
	}
	c.HttpHost = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "HTTP_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | HTTP_PORT is required")
	}
	c.HttpPort = string(pair.Value)

	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "METRIC_HTTP_PORT"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get METRIC_HTTP_PORT from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | METRIC_HTTP_PORT is required")
	}
	c.MetricHttpPort = string(pair.Value)

	// Access Token Config
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "JWT_ACCESS_TOKEN_SECRET"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_SECRET from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_SECRET is required")
	}
	c.JwtAccessTokenSecret = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "JWT_ACCESS_TOKEN_EXPIRATION_TIME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_ACCESS_TOKEN_EXPIRATION_TIME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_ACCESS_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtAccessTokenExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig | JWT_ACCESS_TOKEN_EXPIRATION_TIME is invalid")
	}

	// Refresh Token Token Config
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "JWT_REFRESH_TOKEN_SECRET"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_SECRET from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_SECRET is required")
	}
	c.JwtRefreshTokenSecret = string(pair.Value)
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "JWT_REFRESH_TOKEN_EXPIRATION_TIME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get JWT_REFRESH_TOKEN_EXPIRATION_TIME from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | JWT_REFRESH_TOKEN_EXPIRATION_TIME is required")
	}
	c.JwtRefreshTokenExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig | JWT_REFRESH_TOKEN_EXPIRATION_TIME is invalid")
	}

	// Verification User Login Url Config
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "VERIFICATION_USER_LOGIN_URL"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get VERIFICATION_USER_LOGIN_URL from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul | VERIFICATION_USER_LOGIN_URL is required")
	}
	c.VerificationUserLoginUrl = string(pair.Value)

	// OTP Expiration Time
	pair, _, err = kv.Get(fmt.Sprintf(c.keyPrefix, env, "OTP_EXPIRATION_TIME"), nil)
	if err != nil {
		log.Fatalf("SetConfig | could not get  from consul: %v", err)
	}
	if pair == nil {
		log.Fatal("SetConfig | Consul |  is required")
	}
	c.OtpExpirationTime, err = str2duration.ParseDuration(string(pair.Value))
	if err != nil {
		log.Fatalf("SetConfig |  is invalid")
	}

	return c
}
