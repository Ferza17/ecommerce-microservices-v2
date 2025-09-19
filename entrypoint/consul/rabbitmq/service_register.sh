#!/bin/sh

echo "Registering RabbitMQ to Consul ..."



# Wait for MongoDB to be available (max 10 retries)
echo "Waiting for RabbitMQ to be available..."
max_retries=10
count=0

until nc -z rabbitmq-local 5672 >/dev/null 2>&1; do
  count=$((count+1))
  echo "RabbitMQ not ready yet, attempt $count/$max_retries..."

  if [ $count -ge $max_retries ]; then
    echo "⚠️ RabbitMQ still not available after $max_retries attempts, continuing anyway..."
    break
  fi
  sleep 2
done
echo "RabbitMQ check finished"

# Register MongoDB service
consul services register \
-name=rabbitmq \
-id=rabbitmq-main \
-port=5672 \
-address=rabbitmq-local \
-tag=broker \
-tag=proxy \
-tag=rabbitmq


# Manual health check registration via HTTP API
echo "Adding health checks..."

# Health check for RabbitMQ
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "rabbitmq-Health",
"Name": "RabbitMQ Health Check",
"TCP": "rabbitmq-local:5672",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "rabbitmq-main"
}'

# Verify registration
echo "✅ RabbitMQ registration completed"
