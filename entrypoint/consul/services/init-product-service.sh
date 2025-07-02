#!/bin/sh

initialize_product_service() {
  echo "INIT PRODUCT SERVICE"

  ## Local
  curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/product/ENV
  curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/local/services/product/SERVICE_NAME
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/product/RPC_HOST
  curl --request PUT --data '50055' http://localhost:8500/v1/kv/local/services/product/RPC_PORT
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/product/HTTP_HOST
  curl --request PUT --data '40055' http://localhost:8500/v1/kv/local/services/product/HTTP_PORT
  ## Production
  curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/product/ENV
  curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/production/services/product/SERVICE_NAME
  curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/production/services/product/RPC_HOST
  curl --request PUT --data '50055' http://localhost:8500/v1/kv/production/services/product/RPC_PORT
  curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/production/services/product/HTTP_HOST
  curl --request PUT --data '40055' http://localhost:8500/v1/kv/production/services/product/HTTP_PORT
  echo "DONE INIT PRODUCT SERVICE"
}