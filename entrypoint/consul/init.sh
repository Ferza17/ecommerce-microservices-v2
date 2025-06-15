#!/bin/sh

echo "⏳ Waiting for Consul..."
until curl -s http://localhost:8500/v1/status/leader | grep -q '"'; do
  sleep 1
done


##########################################################################################################################################################################################################################

echo "INIT CONFIG RABBITMQ"
## Local
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/local/broker/rabbitmq/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/local/broker/rabbitmq/RABBITMQ_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/broker/rabbitmq/RABBITMQ_HOST
curl --request PUT --data '5672' http://localhost:8500/v1/kv/local/broker/rabbitmq/RABBITMQ_PORT
## Production
curl --request PUT --data 'rabbitmq' http://localhost:8500/v1/kv/production/broker/rabbitmq/RABBITMQ_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/production/broker/rabbitmq/RABBITMQ_PASSWORD
curl --request PUT --data 'rabbitmq-local' http://localhost:8500/v1/kv/production/broker/rabbitmq/RABBITMQ_HOST
curl --request PUT --data '5672' http://localhost:8500/v1/kv/production/broker/rabbitmq/RABBITMQ_PORT

# EXCHANGE LOCAL
curl --request PUT --data 'product.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/PRODUCT
curl --request PUT --data 'commerce.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/COMMERCE
curl --request PUT --data 'event.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/EVENT
curl --request PUT --data 'notification.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/NOTIFICATION
curl --request PUT --data 'user.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/USER
curl --request PUT --data 'payment.exchange' http://localhost:8500/v1/kv/local/broker/rabbitmq/EXCHANGE/PAYMENT

# EXCHANGE PRODUCTION
curl --request PUT --data 'product.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/PRODUCT
curl --request PUT --data 'commerce.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/COMMERCE
curl --request PUT --data 'event.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/EVENT
curl --request PUT --data 'notification.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/NOTIFICATION
curl --request PUT --data 'user.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/USER
curl --request PUT --data 'payment.exchange' http://localhost:8500/v1/kv/production/broker/rabbitmq/EXCHANGE/PAYMENT

# QUEUE LOCAL USER
curl --request PUT --data 'user.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/USER/CREATED
curl --request PUT --data 'user.updated' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/USER/UPDATED
curl --request PUT --data 'user.login' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/USER/LOGIN
curl --request PUT --data 'user.logout' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/USER/LOGOUT


#QUEUE PRODUCTION USER
curl --request PUT --data 'user.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/USER/CREATED
curl --request PUT --data 'user.updated' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/USER/UPDATED
curl --request PUT --data 'user.login' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/USER/LOGIN
curl --request PUT --data 'user.logout' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/USER/LOGOUT


# QUEUE LOCAL PRODUCT
curl --request PUT --data 'product.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/PRODUCT/CREATED
curl --request PUT --data 'product.updated' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/PRODUCT/UPDATED
curl --request PUT --data 'product.deleted' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/PRODUCT/DELETED

# QUEUE PRODUCTION PRODUCT
curl --request PUT --data 'product.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/PRODUCT/CREATED
curl --request PUT --data 'product.updated' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/PRODUCT/UPDATED
curl --request PUT --data 'product.deleted' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/PRODUCT/DELETED

# QUEUE LOCAL NOTIFICATION
curl --request PUT --data 'notification.email.otp.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED

# QUEUE PRODUCTION NOTIFICATION
curl --request PUT --data 'notification.email.otp.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/NOTIFICATION/EMAIL/OTP/CREATED

# QUEUE LOCAL EVENT
curl --request PUT --data 'event.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/EVENT/CREATED

# QUEUE PRODUCTION EVENT
curl --request PUT --data 'event.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/EVENT/CREATED

# QUEUE LOCAL COMMERCE
curl --request PUT --data 'cart.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED
curl --request PUT --data 'cart.updated' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED
curl --request PUT --data 'cart.deleted' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED

# QUEUE PRODUCTION COMMERCE
curl --request PUT --data 'cart.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/COMMERCE/CART/CREATED
curl --request PUT --data 'cart.updated' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/COMMERCE/CART/UPDATED
curl --request PUT --data 'cart.deleted' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/COMMERCE/CART/DELETED

# QUEUE LOCAL PAYMENT
curl --request PUT --data 'payment.order.created' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED
curl --request PUT --data 'payment.order.delayed.cancelled' http://localhost:8500/v1/kv/local/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED


# QUEUE PRODUCTION PAYMENT
curl --request PUT --data 'payment.order.created' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CREATED
curl --request PUT --data 'payment.order.delayed.cancelled' http://localhost:8500/v1/kv/production/broker/rabbitmq/QUEUE/PAYMENT/ORDER/CANCELLED


##########################################################################################################################################################################################################################

echo "INIT CONFIG DATABASE"
### MongoDB Local
curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_HOST
curl --request PUT --data '27017' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_PORT
curl --request PUT --data 'event-store' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE
curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/COMMERCE
curl --request PUT --data 'notification' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION
## MongoDB Production
curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_PASSWORD
curl --request PUT --data 'mongo-local' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_HOST
curl --request PUT --data '27017' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_PORT
curl --request PUT --data 'event-store' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE
curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/COMMERCE
curl --request PUT --data 'notification' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION

## Postgres Local
curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_PASSWORD
curl --request PUT --data 'disable' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_SSL_MODE
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_HOST
curl --request PUT --data '5432' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_PORT
curl --request PUT --data 'products' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS
curl --request PUT --data 'users' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/USERS
curl --request PUT --data 'payments' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS

## Postgres Production
curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_USERNAME
curl --request PUT --data '1234' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_PASSWORD
curl --request PUT --data 'disable' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_SSL_MODE
curl --request PUT --data 'postgres-local' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_HOST
curl --request PUT --data '5432' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_PORT
curl --request PUT --data 'products' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS
curl --request PUT --data 'users' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/USERS
curl --request PUT --data 'payments' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS

## Elasticsearch Local
curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_HOST
curl --request PUT --data '9200' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_PORT
## Elasticsearch Production
curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_PASSWORD
curl --request PUT --data 'elasticsearch-local' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_HOST
curl --request PUT --data '9200' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_PORT

## Redis Local
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/redis/REDIS_HOST
curl --request PUT --data '6379' http://localhost:8500/v1/kv/local/database/redis/REDIS_PORT
curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/redis/REDIS_PASSWORD
curl --request PUT --data '0' http://localhost:8500/v1/kv/local/database/redis/REDIS_DB
## Redis Production
curl --request PUT --data 'redis-local' http://localhost:8500/v1/kv/production/database/redis/REDIS_HOST
curl --request PUT --data '6379' http://localhost:8500/v1/kv/production/database/redis/REDIS_PORT
curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/redis/REDIS_PASSWORD
curl --request PUT --data '0' http://localhost:8500/v1/kv/production/database/redis/REDIS_DB

##########################################################################################################################################################################################################################

echo "INIT CONFIG SMTP"
## Local
curl --request PUT --data 'ecommerce@email.com' http://localhost:8500/v1/kv/local/smtp/SMTP_SENDER_EMAIL
curl --request PUT --data '' http://localhost:8500/v1/kv/local/smtp/SMTP_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/local/smtp/SMTP_PASSWORD
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/smtp/SMTP_HOST
curl --request PUT --data '1025' http://localhost:8500/v1/kv/local/smtp/SMTP_PORT
## Production
curl --request PUT --data 'ecommerce@email.com' http://localhost:8500/v1/kv/production/smtp/SMTP_SENDER_EMAIL
curl --request PUT --data '' http://localhost:8500/v1/kv/production/smtp/SMTP_USERNAME
curl --request PUT --data '' http://localhost:8500/v1/kv/production/smtp/SMTP_PASSWORD
curl --request PUT --data 'mailhog-local' http://localhost:8500/v1/kv/production/smtp/SMTP_HOST
curl --request PUT --data '1025' http://localhost:8500/v1/kv/production/smtp/SMTP_PORT

##########################################################################################################################################################################################################################

echo "INIT CONFIG TELEMETRY"
# Jaeger Local
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/telemetry/jaeger/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/local/telemetry/jaeger/JAEGER_TELEMETRY_PORT
# Jaeger Production
curl --request PUT --data 'jaeger-local' http://localhost:8500/v1/kv/production/telemetry/jaeger/JAEGER_TELEMETRY_HOST
curl --request PUT --data '14268' http://localhost:8500/v1/kv/production/telemetry/jaeger/JAEGER_TELEMETRY_PORT
##########################################################################################################################################################################################################################

echo "INIT CONFIG API-GATEWAY-SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/api-gateway/ENV
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/local/services/api-gateway/SERVICE_NAME
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/api-gateway/HTTP_HOST
curl --request PUT --data '3000' http://localhost:8500/v1/kv/local/services/api-gateway/HTTP_PORT
## Production
curl --request PUT --data 'local' http://localhost:8500/v1/kv/production/services/api-gateway/ENV
curl --request PUT --data 'api-gateway-service' http://localhost:8500/v1/kv/production/services/api-gateway/SERVICE_NAME
curl --request PUT --data 'api-gateway' http://localhost:8500/v1/kv/production/services/api-gateway/HTTP_HOST
curl --request PUT --data '3000' http://localhost:8500/v1/kv/production/services/api-gateway/HTTP_PORT

##########################################################################################################################################################################################################################
echo "INIT CONFIG COMMERCE-SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/commerce/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/local/services/commerce/SERVICE_NAME
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/commerce/RPC_HOST
curl --request PUT --data '50054' http://localhost:8500/v1/kv/local/services/commerce/RPC_PORT
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/commerce/ENV
curl --request PUT --data 'commerce-service' http://localhost:8500/v1/kv/production/services/commerce/SERVICE_NAME
curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/production/services/commerce/RPC_HOST
curl --request PUT --data '50054' http://localhost:8500/v1/kv/production/services/commerce/RPC_PORT

##########################################################################################################################################################################################################################
echo "INIT EVENT-STORE SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/event-store/ENV
curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/local/services/event-store/SERVICE_NAME
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/event-store/ENV
curl --request PUT --data 'event-store-service' http://localhost:8500/v1/kv/production/services/event-store/SERVICE_NAME

##########################################################################################################################################################################################################################
echo "INIT NOTIFICATION SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/notification/ENV
curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/local/services/notification/SERVICE_NAME
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/notification/ENV
curl --request PUT --data 'notification-service' http://localhost:8500/v1/kv/production/services/notification/SERVICE_NAME

##########################################################################################################################################################################################################################
echo "INIT PRODUCT SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/product/ENV
curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/local/services/product/SERVICE_NAME
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/product/RPC_HOST
curl --request PUT --data '50051' http://localhost:8500/v1/kv/local/services/product/RPC_PORT
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/product/ENV
curl --request PUT --data 'product-service' http://localhost:8500/v1/kv/production/services/product/SERVICE_NAME
curl --request PUT --data 'product' http://localhost:8500/v1/kv/production/services/product/RPC_HOST
curl --request PUT --data '50051' http://localhost:8500/v1/kv/production/services/product/RPC_PORT

##########################################################################################################################################################################################################################
echo "INIT USER SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/user/ENV
curl --request PUT --data 'user-service' http://localhost:8500/v1/kv/local/services/user/SERVICE_NAME
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/user/RPC_HOST
curl --request PUT --data '50052' http://localhost:8500/v1/kv/local/services/user/RPC_PORT
curl --request PUT --data 'ecommerce-service-v2' http://localhost:8500/v1/kv/local/services/user/JWT_ACCESS_TOKEN_SECRET
curl --request PUT --data '1h' http://localhost:8500/v1/kv/local/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME
curl --request PUT --data 'v2-service-ecommerce' http://localhost:8500/v1/kv/local/services/user/JWT_REFRESH_TOKEN_SECRET
curl --request PUT --data '2d' http://localhost:8500/v1/kv/local/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME
curl --request PUT --data '10m' http://localhost:8500/v1/kv/local/services/user/OTP_EXPIRATION_TIME
curl --request PUT --data 'http://localhost:4000?access_token=%s?refresh_token=?' http://localhost:8500/v1/kv/local/services/user/VERIFICATION_USER_LOGIN_URL
## Production
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/user/ENV
curl --request PUT --data 'user-service' http://localhost:8500/v1/kv/production/services/user/SERVICE_NAME
curl --request PUT --data 'user' http://localhost:8500/v1/kv/production/services/user/RPC_HOST
curl --request PUT --data '50052' http://localhost:8500/v1/kv/production/services/user/RPC_PORT
curl --request PUT --data 'ecommerce-service-v2' http://localhost:8500/v1/kv/production/services/user/JWT_ACCESS_TOKEN_SECRET
curl --request PUT --data '1h' http://localhost:8500/v1/kv/production/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME
curl --request PUT --data 'v2-service-ecommerce' http://localhost:8500/v1/kv/production/services/user/JWT_REFRESH_TOKEN_SECRET
curl --request PUT --data '2d' http://localhost:8500/v1/kv/production/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME
curl --request PUT --data '1m' http://localhost:8500/v1/kv/production/services/user/OTP_EXPIRATION_TIME
curl --request PUT --data 'http://localhost:4000?access_token=%s?refresh_token=?' http://localhost:8500/v1/kv/production/services/user/VERIFICATION_USER_LOGIN_URL

##########################################################################################################################################################################################################################
echo "INIT PAYMENT SERVICE"
## Local
curl --request PUT --data 'local' http://localhost:8500/v1/kv/local/services/payment/ENV
curl --request PUT --data 'payment-service' http://localhost:8500/v1/kv/local/services/payment/SERVICE_NAME
curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/services/payment/RPC_HOST
curl --request PUT --data '50055' http://localhost:8500/v1/kv/local/services/payment/RPC_PORT

## Local
curl --request PUT --data 'production' http://localhost:8500/v1/kv/production/services/payment/ENV
curl --request PUT --data 'payment-service' http://localhost:8500/v1/kv/production/services/payment/SERVICE_NAME## Local
curl --request PUT --data 'payment' http://localhost:8500/v1/kv/production/services/payment/RPC_HOST
curl --request PUT --data '50055' http://localhost:8500/v1/kv/production/services/payment/RPC_PORT
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
