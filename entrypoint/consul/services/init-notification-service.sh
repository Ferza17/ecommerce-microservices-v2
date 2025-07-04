echo "INIT NOTIFICATION SERVICE"
## Local
consul kv put local/services/notification/ENV 'local'
consul kv put local/services/notification/SERVICE_NAME 'notification-service'
consul kv put local/services/notification/RPC_HOST '127.0.0.1'
consul kv put local/services/notification/RPC_PORT '50053'
consul kv put local/services/notification/HTTP_HOST 'localhost'
consul kv put local/services/notification/HTTP_PORT '40053'
consul kv put local/services/notification/METRIC_HTTP_PORT '30053'
## Production
consul kv put production/services/notification/ENV 'production'
consul kv put production/services/notification/SERVICE_NAME 'notification-service'
consul kv put production/services/notification/RPC_HOST 'notification-service'
consul kv put production/services/notification/RPC_PORT '50053'
consul kv put production/services/notification/HTTP_HOST 'notification-service'
consul kv put production/services/notification/HTTP_PORT '40053'
consul kv put production/services/notification/METRIC_HTTP_PORT '30053'echo "DONE INIT NOTIFICATION SERVICE"
