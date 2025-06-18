#!/bin/sh

initialize_notification_service() {
  echo "INIT NOTIFICATION SERVICE"

  ## Local
  curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/notification/ENV
  curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/local/services/notification/SERVICE_NAME
  curl --request PUT --data '127.0.0.1' http://localhost:8500/v1/kv/local/services/notification/RPC_HOST
  curl --request PUT --data '50053' http://localhost:8500/v1/kv/local/services/notification/RPC_PORT
  ## Production
  curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/notification/ENV
  curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/production/services/notification/SERVICE_NAME
  curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/production/services/notification/RPC_HOST
  curl --request PUT --data '50053' http://localhost:8500/v1/kv/production/services/notification/RPC_PORT

  echo "DONE INIT NOTIFICATION SERVICE"
}