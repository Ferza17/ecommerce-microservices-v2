echo "Registering Traefik Service in Consul ..."
# Register Traefik Telemetry service
consul services register \
-name=traefik \
-id=traefik-main \
-port=8080 \
-port=80 \
-port=443 \
-port=9000 \
-address=traefik-local \
-tag=traefik \
-tag=reverse-proxy \
-tag=http \
-tag=grpc# Health check for Traefik metric Collector
curl -s -X PUT http://consul-local:8500/v1/agent/check/register \
-H "Content-Type: application/json" \
-d '{
"ID": "traefik-collector-health",
"Name": "traefik Consul Register",
"HTTP": "http://traefik-local:8080/",
"Interval": "30s",
"Timeout": "3s",
"ServiceID": "traefik-main"
}'
# Verify registration
echo "✅ Traefik Consul : registration completed"
