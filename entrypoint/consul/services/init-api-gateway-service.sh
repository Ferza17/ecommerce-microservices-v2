#!/bin/sh


echo "INIT CONFIG API GATEWAY"## Local
consul kv put local/services/api-gateway/ENV 'local'
consul kv put local/services/api-gateway/SERVICE_NAME 'api-gateway-service'
consul kv put local/services/api-gateway/HTTP_HOST '127.0.0.1'
consul kv put local/services/api-gateway/HTTP_PORT '4000'
consul kv put local/services/api-gateway/METRIC_HTTP_PORT '40001'

## Production
consul kv put local/services/api-gateway/ENV 'production'
consul kv put local/services/api-gateway/SERVICE_NAME 'api-gateway-service'
consul kv put local/services/api-gateway/HTTP_HOST '127.0.0.1'
consul kv put local/services/api-gateway/HTTP_PORT '4000'
consul kv put local/services/api-gateway/METRIC_HTTP_PORT '40001'
echo "DONE INIT CONFIG API GATEWAY"
