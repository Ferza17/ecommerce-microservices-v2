#!/bin/sh

initialize_notification_service() {
  echo "INIT NOTIFICATION SERVICE"

  ## Local
  curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/notification/ENV
  curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/local/services/notification/SERVICE_NAME
  ## Production
  curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/notification/ENV
  curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/production/services/notification/SERVICE_NAME

  echo "DONE INIT NOTIFICATION SERVICE"
}