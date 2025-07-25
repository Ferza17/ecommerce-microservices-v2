#!/bin/sh

echo "Registering MongoDB as proxy service in Consul ..."

# Wait for MongoDB to be available
echo "Waiting for MongoDB to be available..."
until nc -z redis-local 6379; do
echo "MongoDB not ready yet, waiting..."
sleep 2
done
echo "MongoDB is available"

# Register MongoDB service
consul services register \
-name=mongodb \
-id=mongodb-main \
-port=6379 \
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
echo "✅ MongoDB proxy registration completed"
