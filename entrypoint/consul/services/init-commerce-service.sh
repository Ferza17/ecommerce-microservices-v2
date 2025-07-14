#!/bin/sh

echo "INIT CONFIG COMMERCE-SERVICE"
## Local
consul kv put local/services/commerce/ENV 'local'
consul kv put local/services/commerce/SERVICE_NAME 'commerce-service'
consul kv put local/services/commerce/RPC_HOST '127.0.0.1'
consul kv put local/services/commerce/RPC_PORT '50051'
consul kv put local/services/commerce/HTTP_HOST '127.0.0.1'
consul kv put local/services/commerce/HTTP_PORT '40051'
consul kv put local/services/commerce/METRIC_HTTP_PORT '30051'
## Production
consul kv put production/services/commerce/ENV 'local'
consul kv put production/services/commerce/SERVICE_NAME 'commerce-service'
consul kv put production/services/commerce/RPC_HOST 'commerce-service'
consul kv put production/services/commerce/RPC_PORT '50051'
consul kv put production/services/commerce/HTTP_HOST 'commerce-service'
consul kv put production/services/commerce/HTTP_PORT '40051'
consul kv put production/services/commerce/METRIC_HTTP_PORT '30051'
echo "DONE INIT CONFIG COMMERCE-SERVICE"
