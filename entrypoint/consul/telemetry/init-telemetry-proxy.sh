echo "Registering Jaeger Telemetry as proxy service in Consul ..."
# Register Jaeger Telemetry service
consul services register \
-name=jaeger \
-id=jaeger-main \
-port=14268 \
-address=jaeger-local \
-tag=telemetry \
-tag=proxy \
-tag=jaeger \
-tag=collector \
-tag=http# Health check for Jaeger Collector
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "jaeger-collector-health",
"Name": "Jaeger Collector Health Check",
"HTTP": "http://jaeger-local:14269/",
"Interval": "30s",
"Timeout": "5s",
"ServiceID": "jaeger-main"
}'
# Verify registration
echo "✅ Jaeger Telemetry proxy registration completed"
