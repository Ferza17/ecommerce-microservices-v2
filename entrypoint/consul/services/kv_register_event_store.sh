#!/bin/sh

echo "INIT EVENT-STORE SERVICE"## Local

consul kv put local/services/event-store/ENV 'local'
consul kv put local/services/event-store/SERVICE_NAME 'event-store-service'
consul kv put local/services/event-store/RPC_HOST '127.0.0.1'
consul kv put local/services/event-store/RPC_PORT '50052'
consul kv put local/services/event-store/HTTP_HOST '127.0.0.1'
consul kv put local/services/event-store/HTTP_PORT '40052'
consul kv put local/services/event-store/METRIC_HTTP_PORT '30052'

## Production

consul kv put production/services/event-store/ENV 'local'
consul kv put production/services/event-store/SERVICE_NAME 'event-store-service'
consul kv put production/services/event-store/RPC_HOST 'event-store-service'
consul kv put production/services/event-store/RPC_PORT '50052'
consul kv put production/services/event-store/HTTP_HOST 'event-store-service'
consul kv put production/services/event-store/HTTP_PORT '40052'
consul kv put production/services/event-store/METRIC_HTTP_PORT '30052'

echo "DONE INIT EVENT-STORE SERVICE"
