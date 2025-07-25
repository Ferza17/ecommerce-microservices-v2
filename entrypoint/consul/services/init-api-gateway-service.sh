#!/bin/sh


echo "INIT CONFIG API GATEWAY"## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/api-gateway/ENV
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/local/services/api-gateway/SERVICE_NAME
curl --request PUT --data '127.0.0.1' http://localhost:8500/v1/kv/local/services/api-gateway/HTTP_HOST
curl --request PUT --data '3000' http://localhost:8500/v1/kv/local/services/api-gateway/HTTP_PORT
curl --request PUT --data '127.0.0.1' http://localhost:8500/v1/kv/local/services/api-gateway/RPC_HOST
curl --request PUT --data '50000' http://localhost:8500/v1/kv/local/services/api-gateway/RPC_PORT
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/api-gateway/ENV
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/production/services/api-gateway/SERVICE_NAME
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/production/services/api-gateway/HTTP_HOST
curl --request PUT --data '3000' http://localhost:8500/v1/kv/production/services/api-gateway/HTTP_PORT
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/production/services/api-gateway/RPC_HOST
curl --request PUT --data '50000' http://localhost:8500/v1/kv/production/services/api-gateway/RPC_PORT
echo "DONE INIT CONFIG API GATEWAY"
