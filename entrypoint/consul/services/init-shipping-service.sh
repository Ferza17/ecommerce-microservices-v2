#!/bin/sh


echo "INIT SHIPPING service"

## LOCAL
consul kv put local/services/shipping/ENV 'local'
consul kv put local/services/shipping/SERVICE_NAME 'shipping-services'
consul kv put local/services/shipping/RPC_HOST '127.0.0.1'
consul kv put local/services/shipping/RPC_PORT '50057'
consul kv put local/services/shipping/HTTP_HOST '127.0.0.1'
consul kv put local/services/shipping/HTTP_PORT '40057'
consul kv put local/services/shipping/METRIC_HTTP_PORT '30057'

## PRODUCTION
consul kv put production/services/shipping/ENV 'production'
consul kv put production/services/shipping/SERVICE_NAME 'shipping-services'
consul kv put production/services/shipping/RPC_HOST 'shipping-services'
consul kv put production/services/shipping/RPC_PORT '50057'
consul kv put production/services/shipping/HTTP_HOST 'shipping-services'
consul kv put production/services/shipping/HTTP_PORT '40057'
consul kv put production/services/shipping/METRIC_HTTP_PORT '30057'
echo "DONE INIT SHIPPING service"






