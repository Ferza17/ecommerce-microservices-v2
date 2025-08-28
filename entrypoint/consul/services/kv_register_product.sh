#!/bin/sh


echo "INIT PRODUCT SERVICE"
## Local
consul kv put local/services/product/ENV 'local'
consul kv put local/services/product/SERVICE_NAME 'product-service'
consul kv put local/services/product/RPC_HOST 'localhost'
consul kv put local/services/product/RPC_PORT '50055'
consul kv put local/services/product/HTTP_HOST 'localhost'
consul kv put local/services/product/HTTP_PORT '40055'
consul kv put local/services/product/METRIC_HTTP_PORT '30055'
## Production
consul kv put production/services/product/ENV 'production'
consul kv put production/services/product/SERVICE_NAME 'product-service'
consul kv put production/services/product/RPC_HOST 'product-service'
consul kv put production/services/product/RPC_PORT '50055'
consul kv put production/services/product/HTTP_HOST 'product-service'
consul kv put production/services/product/HTTP_PORT '40055'
consul kv put production/services/product/METRIC_HTTP_PORT '30055'
echo "DONE INIT PRODUCT SERVICE"
