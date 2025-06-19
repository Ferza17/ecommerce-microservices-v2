#!/bin/sh

initialize_telemetry(){
  echo "INIT CONFIG TELEMETRY"
  # Jaeger Local
  consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_HOST "localhost"
  consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_PORT ""

  # Jaeger Production
  consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_HOST "jaeger-local"
  consul kv put local/telemetry/jaeger/JAEGER_TELEMETRY_PORT "14268"

  echo "DONE INIT CONFIG TELEMETRY"
}

initialize_telemetry