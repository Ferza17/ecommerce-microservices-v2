#!/bin/sh

initialize_rabbitmq () {
  echo "INIT CONFIG RABBITMQ"
  ## Local

  consul kv put local/broker/rabbitmq/RABBITMQ_USERNAME "rabbitmq"
  consul kv put local/broker/rabbitmq/RABBITMQ_PASSWORD "1234"
  consul kv put local/broker/rabbitmq/RABBITMQ_HOST "localhost"
  consul kv put local/broker/rabbitmq/RABBITMQ_PORT "5672"
  ## Production
  consul kv put production/broker/rabbitmq/RABBITMQ_USERNAME "rabbitmq"
  consul kv put production/broker/rabbitmq/RABBITMQ_PASSWORD "1234"
  consul kv put production/broker/rabbitmq/RABBITMQ_HOST "rabbitmq-local"
  consul kv put production/broker/rabbitmq/RABBITMQ_PORT "5672"


  echo "DONE INIT CONFIG RABBITMQ"
}

initialize_rabbitmq