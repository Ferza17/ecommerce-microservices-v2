#!/bin/sh

initialize_telemetry(){
  echo "INIT CONFIG TELEMETRY"
  # Jaeger Local
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/telemetry/jaeger/JAEGER_TELEMETRY_HOST
  curl --request PUT --data '14268' http://localhost:8500/v1/kv/local/telemetry/jaeger/JAEGER_TELEMETRY_PORT
  # Jaeger Production
  curl --request PUT --data 'jaeger-local' http://localhost:8500/v1/kv/production/telemetry/jaeger/JAEGER_TELEMETRY_HOST
  curl --request PUT --data '14268' http://localhost:8500/v1/kv/production/telemetry/jaeger/JAEGER_TELEMETRY_PORT

  echo "DONE INIT CONFIG TELEMETRY"
}