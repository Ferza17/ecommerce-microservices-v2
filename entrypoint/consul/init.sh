#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://localhost:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done


# LOAD CONFIG RABBITMQ
##########################################################################################################################################################################################################################
source /rabbitmq/init-rabbitmq.sh
initialize_rabbitmq

source /rabbitmq/init-rabbitmq-exchange.sh
initialize_rabbitmq_exchange

source /rabbitmq/init-rabbitmq-queue.sh
initialize_rabbitmq_queue

# LOAD CONFIG DATABASE
##########################################################################################################################################################################################################################
source /database/init-database-elasticsearch.sh
initialize_database_elasticsearch

source /database/init-database-mongodb.sh
initialize_database_mongodb

source /database/init-database-postgresql.sh
initialize_database_postgresql

source /database/init-database-redis.sh
initialize_database_redis

# LOAD CONFIG SMTP
##########################################################################################################################################################################################################################
source /smtp/init-smtp.sh
initialize_smtp

# LOAD CONFIG TELEMETRY
##########################################################################################################################################################################################################################
source /telemetry/init-telemetry.sh
initialize_telemetry

# LOAD CONFIG SERVICES
##########################################################################################################################################################################################################################
source /services/init-api-gateway-service.sh
initialize_service_api_gateway

source /services/init-commerce-service.sh
initialize_commerce_service

source /services/init-event-store-service.sh
initialize_event_store_service

source /services/init-notification-service.sh
initialize_notification_service

source /services/init-payment-service.sh
initialize_payment_service

source /services/init-product-service.sh
initialize_product_service

source /services/init-user-service.sh
initialize_user_service

##########################################################################################################################################################################################################################

echo "INIT CONFIG COMMON"
## Local
curl --request PUT --data 'PENDING' http://localhost:8500/v1/kv/local/common/SAGA_STATUS/PENDING
curl --request PUT --data 'SUCCESS' http://localhost:8500/v1/kv/local/common/SAGA_STATUS/SUCCESS
curl --request PUT --data 'FAILED' http://localhost:8500/v1/kv/local/common/SAGA_STATUS/FAILED

## Production
curl --request PUT --data 'PENDING' http://localhost:8500/v1/kv/production/common/SAGA_STATUS/PENDING
curl --request PUT --data 'SUCCESS' http://localhost:8500/v1/kv/production/common/SAGA_STATUS/SUCCESS
curl --request PUT --data 'FAILED' http://localhost:8500/v1/kv/production/common/SAGA_STATUS/FAILED

echo "✅ Done setting key-values."
