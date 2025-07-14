#!/bin/sh

echo "INIT CONFIG TELEMETRY"
# Jaeger Local
consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_HOST "localhost"
consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_PORT "14268"
# Jaeger Production
consul kv put production/telemetry/jaeger/JAEGER_TELEMETRY_HOST "jaeger-local"
consul kv put production/telemetry/jaeger/JAEGER_TELEMETRY_PORT "14268"
echo "DONE INIT CONFIG TELEMETRY"
