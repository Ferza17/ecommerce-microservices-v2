#!/bin/bash

echo "Registering Elasticsearch to Consul ..."

# Wait for Elasticsearch to be available (max 10 retries)
echo "Waiting for Elasticsearch to be available..."
max_retries=10
count=0

until nc -z elasticsearch-local 9200 >/dev/null 2>&1; do
  count=$((count+1))
  echo "Elasticsearch not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ Elasticsearch still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "Elasticsearch check finished"

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
