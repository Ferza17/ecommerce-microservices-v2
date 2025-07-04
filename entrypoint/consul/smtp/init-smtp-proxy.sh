echo "Registering Mailhog SMTP as proxy service in Consul ..."

# Wait for Mailhog SMTP to be available
echo "Waiting for Mailhog to be available..."
until nc -z mailhog-local 1025; do
echo "Mailhog not ready yet, waiting..."
sleep 2
done
echo "Mailhog is available"


# Register Mailhog SMTP service
consul services register \
-name=mailhog \
-id=mailhog-main \
-port=1025 \
-address=mailhog-local \
-tag=smtp \
-tag=proxy \
-tag=mailhog \


# Health check for Jaeger Collector
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "mailhog-ui-health",
"Name": "Mailhog UI Health Check",
"HTTP": "http://mailhog-local:8025/",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "mailhog-main"
}'
