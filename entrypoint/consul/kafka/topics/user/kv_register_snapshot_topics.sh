#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE USER TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/USER/USER_CREATED "snapshot-users-user_created"
consul kv put local/broker/kafka/TOPICS/USER/CONFIRM/USER_CREATED "confirm-snapshot-users-user_created"
consul kv put local/broker/kafka/TOPICS/USER/COMPENSATE/USER_CREATED "compensate-snapshot-users-user_created"

consul kv put local/broker/kafka/TOPICS/USER/USER_UPDATED "snapshot-users-user_updated"
consul kv put local/broker/kafka/TOPICS/USER/CONFIRM/USER_UPDATED "confirm-snapshot-users-user_updated"
consul kv put local/broker/kafka/TOPICS/USER/COMPENSATE/USER_UPDATED "compensate-snapshot-users-user_updated"

consul kv put local/broker/kafka/TOPICS/USER/USER_LOGIN "snapshot-users-user_login"
consul kv put local/broker/kafka/TOPICS/USER/CONFIRM/USER_LOGIN "confirm-snapshot-users-user_login"
consul kv put local/broker/kafka/TOPICS/USER/COMPENSATE/USER_LOGIN "compensate-snapshot-users-user_login"

consul kv put local/broker/kafka/TOPICS/USER/USER_LOGOUT "snapshot-users-user_logout"
consul kv put local/broker/kafka/TOPICS/USER/CONFIRM/USER_LOGOUT "confirm-snapshot-users-user_logout"
consul kv put local/broker/kafka/TOPICS/USER/COMPENSATE/USER_LOGOUT "compensate-snapshot-users-user_logout"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/USER/USER_CREATED "snapshot-users-user_created"
consul kv put production/broker/kafka/TOPICS/USER/CONFIRM/USER_CREATED "confirm-snapshot-users-user_created"
consul kv put production/broker/kafka/TOPICS/USER/COMPENSATE/USER_CREATED "compensate-snapshot-users-user_created"

consul kv put production/broker/kafka/TOPICS/USER/USER_UPDATED "snapshot-users-user_updated"
consul kv put production/broker/kafka/TOPICS/USER/CONFIRM/USER_UPDATED "confirm-snapshot-users-user_updated"
consul kv put production/broker/kafka/TOPICS/USER/COMPENSATE/USER_UPDATED "compensate-snapshot-users-user_updated"

consul kv put production/broker/kafka/TOPICS/USER/USER_LOGIN "snapshot-users-user_login"
consul kv put production/broker/kafka/TOPICS/USER/CONFIRM/USER_LOGIN "confirm-snapshot-users-user_login"
consul kv put production/broker/kafka/TOPICS/USER/COMPENSATE/USER_LOGIN "compensate-snapshot-users-user_login"

consul kv put production/broker/kafka/TOPICS/USER/USER_LOGOUT "snapshot-users-user_logout"
consul kv put production/broker/kafka/TOPICS/USER/CONFIRM/USER_LOGOUT "confirm-snapshot-users-user_logout"
consul kv put production/broker/kafka/TOPICS/USER/COMPENSATE/USER_LOGOUT "compensate-snapshot-users-user_logout"

echo "DONE INIT CONFIG KAFKA NAMESPACE USER TOPICS"
