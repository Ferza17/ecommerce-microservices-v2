#!/bin/bash

echo "Registering Redis to Consul ..."

# Wait for Redis to be available (max 10 retries)
echo "Waiting for Redis to be available..."
max_retries=10
count=0

until nc -z redis-local 6379 >/dev/null 2>&1; do
  count=$((count+1))
  echo "Redis not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ Redis still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "Redis check finished"

# Register Redis service
consul services register \
-name=redis \
-id=redis-main \
-port=6379 \
-address=redis-local \
-tag=database \
-tag=cache \
-tag=proxy \
-tag=redis

# Manual health check registration via HTTP API
echo "Adding health checks..."

# Health check for redis
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "redis-health",
"Name": "Redis Health Check",
"TCP": "redis-local:6379",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "redis-main"
}'

# Verify registration
echo "✅ Redis proxy registration completed"





