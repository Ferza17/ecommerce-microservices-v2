#!/bin/bash

echo "Registering Elasticsearch as proxy service in Consul ..."

 # Wait for Redis to be available
echo "Waiting for Redis to be available..."
until nc -z elasticsearch-local 9200; do
echo "Elasticsearch not ready yet, waiting..."
sleep 2
done
echo "Elasticsearch is available"

# Register Elasticsearch service
consul services register \
-name=elasticsearch \
-id=elasticsearch-main \
-port=9200 \
-address=elasticsearch \
-tag=database \
-tag=proxy \
-tag=elasticsearch

# Manual health check registration via HTTP API
echo "Adding health checks..."

# Health check for elasticsearch (TCP)
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "elasticsearch-tcp-health",
"Name": "Elasticsearch TCP Health Check",
"TCP": "elasticsearch-local:9200",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "elasticsearch-main"
}'

# Health check for elasticsearch (HTTP)
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "elasticsearch-http-health",
"Name": "Elasticsearch HTTP Health Check",
"HTTP": "http://elasticsearch-local:9200/_cluster/health",
"Method": "GET",
"Interval": "30s",
"Timeout": "10s",
"ServiceID": "elasticsearch-main"
}'


# Verify registration
echo "✅ Elasticsearch proxy registration completed"
