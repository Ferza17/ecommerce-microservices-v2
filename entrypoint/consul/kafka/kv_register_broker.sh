#!/bin/sh

echo "INIT CONFIG KAFKA BROKER"
consul kv put local/broker/kafka/BROKER_1 "localhost:9092"

echo "DONE INIT CONFIG KAFKA BROKER"
consul kv put production/broker/kafka/BROKER_1 "kafka-local-broker-1:9092"
