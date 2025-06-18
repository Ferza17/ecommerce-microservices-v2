#!/bin/sh

initialize_database_redis() {
  echo "INIT CONFIG DATABASE REDIS"

  ## Redis Local
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/redis/REDIS_HOST
  curl --request PUT --data '6379' http://localhost:8500/v1/kv/local/database/redis/REDIS_PORT
  curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/redis/REDIS_PASSWORD
  curl --request PUT --data '0' http://localhost:8500/v1/kv/local/database/redis/REDIS_DB
  ## Redis Production
  curl --request PUT --data 'redis-local' http://localhost:8500/v1/kv/production/database/redis/REDIS_HOST
  curl --request PUT --data '6379' http://localhost:8500/v1/kv/production/database/redis/REDIS_PORT
  curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/redis/REDIS_PASSWORD
  curl --request PUT --data '0' http://localhost:8500/v1/kv/production/database/redis/REDIS_DB

  echo "DONE INIT CONFIG DATABASE REDIS"
}