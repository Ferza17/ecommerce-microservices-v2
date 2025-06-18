#!/bin/sh

initialize_rabbitmq_exchange() {
  echo "INIT CONFIG RABBITMQ EXCHANGE"

  # EXCHANGE LOCAL
  curl --request PUT --data 'product.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/PRODUCT
  curl --request PUT --data 'commerce.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/COMMERCE
  curl --request PUT --data 'event.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/EVENT
  curl --request PUT --data 'notification.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/NOTIFICATION
  curl --request PUT --data 'user.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/USER
  curl --request PUT --data 'payment.direct.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT
  curl --request PUT --data 'payment.delayed.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED


  # EXCHANGE PRODUCTION
  curl --request PUT --data 'product.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/PRODUCT
  curl --request PUT --data 'commerce.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/COMMERCE
  curl --request PUT --data 'event.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/EVENT
  curl --request PUT --data 'notification.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/NOTIFICATION
  curl --request PUT --data 'user.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/USER
  curl --request PUT --data 'payment.direct.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/PAYMENT/DIRECT
  curl --request PUT --data 'payment.delayed.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/PAYMENT/DELAYED

  echo "DONE INIT CONFIG RABBITMQ EXCHANGE"
}