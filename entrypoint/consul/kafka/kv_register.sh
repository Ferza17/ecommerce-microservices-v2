#!/bin/sh



echo "INIT CONFIG KAFKA BROKER"
# BROKER local
consul kv put local/broker/kafka/BROKER_1 "localhost:9092"
# BROKER production
consul kv put production/broker/kafka/BROKER_1 "kafka-local-broker-1:29092"

# SCHEMA REGISTRY local
consul kv put local/broker/kafka/SCHEMA_REGISTRY "http://localhost:8081"
# SCHEMA REGISTRY local
consul kv put production/broker/kafka/SCHEMA_REGISTRY "http://schema-registry-local:8081"

echo "DONE INIT CONFIG KAFKA BROKER"
