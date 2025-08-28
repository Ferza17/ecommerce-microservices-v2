#!/bin/sh

echo "INIT CONSUL KV DATABASE REDIS"

## Redis Local
consul kv put local/database/redis/REDIS_HOST "localhost"
consul kv put local/database/redis/REDIS_PORT "6379"
consul kv put local/database/redis/REDIS_PASSWORD ""
consul kv put local/database/redis/REDIS_DB "0"

## Redis Production
consul kv put production/database/redis/REDIS_HOST "redis-local"
consul kv put production/database/redis/REDIS_PORT "6379"
consul kv put production/database/redis/REDIS_PASSWORD ""
consul kv put production/database/redis/REDIS_DB "0"

echo "✅ DONE INIT CONSUL KV DATABASE REDIS"
