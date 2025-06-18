#!/bin/sh

initialize_payment_service() {
  echo "INIT PAYMENT SERVICE"
  ## Local
  curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/payment/ENV
  curl --request PUT --data 'payment-service' http://localhost:8500/v1/kv/local/services/payment/SERVICE_NAME
  curl --request PUT --data '600000' http://localhost:8500/v1/kv/local/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/payment/RPC_HOST
  curl --request PUT --data '50054' http://localhost:8500/v1/kv/local/services/payment/RPC_PORT

  ## Local
  curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/payment/ENV
  curl --request PUT --data 'payment-service' http://localhost:8500/v1/kv/production/services/payment/SERVICE_NAME
  curl --request PUT --data '600000' http://localhost:8500/v1/kv/production/services/payment/PAYMENT_ORDER_CANCELLED_IN_MS
  curl --request PUT --data 'payment-service' http://localhost:8500/v1/kv/production/services/payment/RPC_HOST
  curl --request PUT --data '50054' http://localhost:8500/v1/kv/production/services/payment/RPC_PORT

  echo "DONE INIT PAYMENT SERVICE"

}