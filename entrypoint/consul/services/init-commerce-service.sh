echo "INIT CONFIG COMMERCE-SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/commerce/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/local/services/commerce/SERVICE_NAME
curl --request PUT --data '127.0.0.1' http://localhost:8500/v1/kv/local/services/commerce/RPC_HOST
curl --request PUT --data '50051' http://localhost:8500/v1/kv/local/services/commerce/RPC_PORT
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/commerce/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/production/services/commerce/SERVICE_NAME
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/production/services/commerce/RPC_HOST
curl --request PUT --data '50051' http://localhost:8500/v1/kv/production/services/commerce/RPC_PORT
echo "DONE INIT CONFIG COMMERCE-SERVICE"
