#!/bin/bash

echo "Registering PostgresSQL to Consul ..."

# Wait for PostgresSQL to be available (max 10 retries)
echo "Waiting for PostgresSQL to be available..."
max_retries=10
count=0

until nc -z postgres-local 5432 >/dev/null 2>&1; do
  count=$((count+1))
  echo "PostgresSQL not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ PostgresSQL still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "PostgresSQL check finished"


# Register MongoDB service
consul services register \
-name=postgresSQL \
-id=postgres-main \
-port=5432 \
-address=postgres-local \
-tag=database \
-tag=cache \
-tag=proxy \
-tag=postgresSQL

# Manual health check registration via HTTP API
echo "Adding health checks..."

# Health check for MongoDB
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "postgresSQL-health",
"Name": "PostgresSQL Health Check",
"TCP": "postgres-local:5432",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "postgres-main"
}'

# Verify registration
echo "✅ PostgresSQL proxy registration completed"
