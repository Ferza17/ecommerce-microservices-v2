#!/bin/sh

echo "INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE USER"
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/USERS "sink-pg-users-users"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/DLQ/USERS "dlq-sink-pg-users-users"

consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/ROLES "sink-pg-users-roles"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/DLQ/ROLES "dlq-sink-pg-users-roles"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/USERS "sink-pg-users-users"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/DLQ/USERS "dlq-sink-pg-users-users"

consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/ROLES "sink-pg-users-roles"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/DLQ/ROLES "dlq-sink-pg-users-roles"

echo "DONE INIT CONFIG KAFKA CONNECTOR TOPICS NAMESPACE USER"
