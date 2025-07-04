

echo "Registering Jaeger Telemetry as proxy service in Consul ..."


# Register Jaeger Telemetry service
consul services register \
-name=prometheus \
-id=prometheus-main \
-port=9090 \
-address=prometheus-local \
-tag=proxy \
-tag=prometheus \
-tag=metric \
-tag=collector \
-tag=http

# Health check for prometheus metric Collector
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "prometheus-collector-health",
"Name": "Prometheus Consul Register",
"HTTP": "http://prometheus-local:9090/",
"Interval": "30s",
"Timeout": "3s",
"ServiceID": "prometheus-main"
}'

# Verify registration
echo "✅ Prometheus Consul : registration completed"
