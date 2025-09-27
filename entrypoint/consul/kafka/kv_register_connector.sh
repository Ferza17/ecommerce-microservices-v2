#!/bin/sh

echo "INIT CONFIG KAFKA TOPICS SINK & SOURCE CONNECTOR"
# ----------- NOTIFICATION

## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/NOTIFICATION-TEMPLATES "source-mongo-notification-notification_templates"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SOURCE/MONGO/NOTIFICATION/NOTIFICATION-TEMPLATES "source-mongo-notification-notification_templates"

# ----------- PAYMENT
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-PROVIDERS "sink-pg-payments-payment_providers"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-ITEMS "sink-pg-payments-payment_items"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENTS "sink-pg-payments-payments"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-PROVIDERS "sink-pg-payments-payment_providers"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENT-ITEMS "sink-pg-payments-payment_items"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PAYMENT/PAYMENTS "sink-pg-payments-payments"

# ----------- PRODUCT
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/ES/PRODUCT/PRODUCTS "sink-es-products-products"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PRODUCT/PRODUCTS "sink-pg-products-products"

#PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/ES/PRODUCT/PRODUCTS "sink-es-products-products"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/PRODUCT/PRODUCTS "sink-pg-products-products"

# ----------- SHIPPING
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPING-PROVIDERS "sink-pg-shippings-shipping_providers"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPINGS "sink-pg-shippings-shippings"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPING-PROVIDERS "sink-pg-shippings-shipping_providers"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/SHIPPING/SHIPPINGS "sink-pg-shippings-shippings"

# ----------- USER
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/USERS "sink-pg-users-users"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/ROLES "sink-pg-users-roles"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/USERS "sink-pg-users-users"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/PG/USER/ROLES "sink-pg-users-roles"

# ----------- EVENT
## LOCAL
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/COMMERCE "sink-mongo-events-commerce_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/NOTIFICATION "sink-mongo-events-notification_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/PAYMENT "sink-mongo-events-payment_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/PRODUCT "sink-mongo-events-product_event_stores"
consul kv put local/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/USER "sink-mongo-events-user_event_stores"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/COMMERCE "sink-mongo-events-commerce_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/NOTIFICATION "sink-mongo-events-notification_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/PAYMENT "sink-mongo-events-payment_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/PRODUCT "sink-mongo-events-product_event_stores"
consul kv put production/broker/kafka/TOPICS/CONNECTOR/SINK/MONGO/EVENTS/USER "sink-mongo-events-user_event_stores"

echo "DONE INIT CONFIG KAFKA TOPICS SINK & SOURCE CONNECTOR"
