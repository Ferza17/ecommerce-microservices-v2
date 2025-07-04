echo "INIT EVENT-STORE SERVICE"## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/event-store/ENV
curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/local/services/event-store/SERVICE_NAME
curl --request PUT --data '127.0.0.1' http://localhost:8500/v1/kv/local/services/event-store/RPC_HOST
curl --request PUT --data '50052' http://localhost:8500/v1/kv/local/services/event-store/RPC_PORT
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/event-store/ENV
curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/production/services/event-store/SERVICE_NAME
curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/production/services/event-store/RPC_HOST
curl --request PUT --data '50052' http://localhost:8500/v1/kv/production/services/event-store/RPC_PORTecho "DONE INIT EVENT-STORE SERVICE"
