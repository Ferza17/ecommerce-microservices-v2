#!/bin/sh

initialize_event_store_service(){
  echo "INIT EVENT-STORE SERVICE"

  ## Local
  curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/event-store/ENV
  curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/local/services/event-store/SERVICE_NAME
  ## Production
  curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/event-store/ENV
  curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/production/services/event-store/SERVICE_NAME

  echo "DONE INIT EVENT-STORE SERVICE"
}