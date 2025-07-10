echo "INIT SHIPPING SERVICE"

## LOCAL
consul kv put local/service/shipping/ENV 'local'
consul kv put local/service/shipping/SERVICE_NAME 'shipping-service'
consul kv put local/service/shipping/RPC_HOST 'localhost'
consul kv put local/service/shipping/RPC_PORT '50057'
consul kv put local/service/shipping/HTTP_HOST 'localhost'
consul kv put local/service/shipping/HTTP_PORT '40057'
consul kv put local/service/shipping/METRIC_HTTP_PORT '30057'

## PRODUCTION
consul kv put production/service/shipping/ENV 'production'
consul kv put production/service/shipping/SERVICE_NAME 'shipping-service'
consul kv put production/service/shipping/RPC_HOST 'shipping-service'
consul kv put production/service/shipping/RPC_PORT '50057'
consul kv put production/service/shipping/HTTP_HOST 'shipping-service'
consul kv put production/service/shipping/HTTP_PORT '40057'
consul kv put production/service/shipping/METRIC_HTTP_PORT '30057'
echo "DONE INIT SHIPPING SERVICE"






