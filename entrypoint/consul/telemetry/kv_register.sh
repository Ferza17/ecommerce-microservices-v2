#!/bin/sh

echo "INIT CONFIG TELEMETRY"
# Jaeger Local
consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_HOST "localhost"
consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_PORT "14268"
consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_RPC_PORT "4317"

# Jaeger Production
consul kv put production/telemetry/jaeger/JAEGER_TELEMETRY_HOST "jaeger-local"
consul kv put production/telemetry/jaeger/JAEGER_TELEMETRY_PORT "14268"
consul kv put production/telemetry/jaeger/JAEGER_TELEMETRY_RPC_PORT "4317"
echo "DONE INIT CONFIG TELEMETRY"
