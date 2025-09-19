#!/bin/sh

echo "Registering MongoDB to Consul ..."

# Wait for MongoDB to be available (max 10 retries)
echo "Waiting for MongoDB to be available..."
max_retries=10
count=0

until nc -z mongo-local 27017 >/dev/null 2>&1; do
  count=$((count+1))
  echo "MongoDB not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ MongoDB still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "MongoDB check finished"

# Register MongoDB service
consul services register \
  -name=mongodb \
  -id=mongodb-main \
  -port=27017 \
  -address=mongo-local \
  -tag=database \
  -tag=cache \
  -tag=proxy \
  -tag=mongodb

# Manual health check registration via HTTP API
echo "Adding health checks..."

# Health check for MongoDB
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
  -H "Content-Type: application/json" \
  -d '{
    "ID": "mongodb-health",
    "Name": "MongoDB Health Check",
    "TCP": "mongo-local:27017",
    "Interval": "30s",
    "Timeout": "5s",
    "ServiceID": "mongodb-main"
  }'

# Verify registration
echo "✅ MONGODB service register is completed"
