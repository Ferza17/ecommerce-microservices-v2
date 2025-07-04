

echo "Registering RabbitMQ as proxy service in Consul ..."

# Wait for RabbitMQ to be available
echo "Waiting for RabbitMQ to be available..."
until nc -z rabbitmq-local 5672; do
echo "RabbitMQ not ready yet, waiting..."
sleep 2
done
echo "RabbitMQ is available"

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
echo "✅ RabbitMQ proxy registration completed"
