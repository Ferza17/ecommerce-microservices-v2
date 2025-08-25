#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://consul-local:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done


# LOAD CONFIG RABBITMQ
##########################################################################################################################################################################################################################
sh ./rabbitmq/init-rabbitmq.sh
sh ./rabbitmq/init-rabbitmq-exchange.sh
sh ./rabbitmq/init-rabbitmq-queue.sh
sh ./rabbitmq/init-rabbitmq-proxy.sh

# LOAD CONFIG DATABASE
##########################################################################################################################################################################################################################
# ELASTICSEARCH
sh ./database/elasticsearch/init-database-elasticsearch.sh
sh ./database/elasticsearch/init-database-elasticsearch-proxy.sh

# POSTGRES
sh ./database/postgres/init-database-postgresql.sh
sh ./database/postgres/init-database-postgresql-proxy.sh

# REDIS
sh ./database/redis/init-database-redis.sh
sh ./database/redis/init-database-redis-proxy.sh

# MONGODB
sh ./database/mongodb/init-database-mongodb.sh
sh ./database/mongodb/init-database-mongodb-proxy.sh


# LOAD CONFIG SMTP
##########################################################################################################################################################################################################################
sh ./smtp/init-smtp.sh
sh ./smtp/init-smtp-proxy.sh

# LOAD CONFIG TELEMETRY
##########################################################################################################################################################################################################################
 sh ./telemetry/init-telemetry.sh
 sh ./telemetry/init-telemetry-proxy.sh

# LOAD CONFIG PROMETHEUS AND METRICS COLLECTOR
##########################################################################################################################################################################################################################
sh ./prometheus/init-prometheus-proxy.sh
sh ./postgres-exporter/register.sh

# LOAD CONFIG OPEN POLICY AGENT
##########################################################################################################################################################################################################################
 sh ./opa/init.sh
# sh ./opa/register.sh

# LOAD CONFIG TRAEFIK
##########################################################################################################################################################################################################################
sh ./traefik/init-traefik-proxy.sh

# LOAD CONFIG SERVICES
##########################################################################################################################################################################################################################
sh ./services/init-api-gateway-service.sh
sh ./services/init-commerce-service.sh
sh ./services/init-event-store-service.sh
sh ./services/init-notification-service.sh
sh ./services/init-payment-service.sh
sh ./services/init-product-service.sh
sh ./services/init-user-service.sh
sh ./services/init-shipping-service.sh

##########################################################################################################################################################################################################################

echo "INIT CONFIG COMMON"
## LOCAL
consul kv put local/common/SAGA_STATUS/PENDING 'PENDING'
consul kv put local/common/SAGA_STATUS/SUCCESS 'SUCCESS'
consul kv put local/common/SAGA_STATUS/FAILED 'FAILED'

## PRODUCTION
consul kv put production/common/SAGA_STATUS/PENDING 'PENDING'
consul kv put production/common/SAGA_STATUS/SUCCESS 'SUCCESS'
consul kv put production/common/SAGA_STATUS/FAILED 'FAILED'

echo "✅ Done setting key-values."
