#!/bin/bash

echo "Registering PostgresSQL as proxy service in Consul ..."

# Wait for MongoDB to be available
echo "Waiting for PostgresSQL to be available..."
until nc -z postgres-local 5432; do
echo "PostgresSQL not ready yet, waiting..."
sleep 2
done
echo "MongoDB is available"


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
