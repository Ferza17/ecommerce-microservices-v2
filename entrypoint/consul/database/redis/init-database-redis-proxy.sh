#!/bin/bash

echo "Registering Redis as proxy service in Consul ..."

# Wait for Redis to be available
echo "Waiting for Redis to be available..."
until nc -z redis-local 6379; do
echo "Redis not ready yet, waiting..."
sleep 2
done
echo "Redis is available"

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





