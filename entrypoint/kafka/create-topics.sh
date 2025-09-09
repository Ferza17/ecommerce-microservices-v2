#!/bin/bash

echo ">>> Creating Kafka topics..."
KAFKA_BROKER="kafka-local:9092"

# Define topics (name:partitions:replication-factor)
TOPICS=(
  "user.user_created:3:1"
  "user.user_updated:3:1"
  "user.user_login:3:1"
  "user.user_logout:1:1"
  "product.product_created:3:1"
  "product.product_updated:3:1"
  "product.product_deleted:3:1"
  "notification.email_otp_created:2:1"
  "notification.email_payment_order_created:2:1"
  "event.event_created:3:1"
  "commerce.cart_created:3:1"
  "commerce.cart_updated:3:1"
  "commerce.cart_deleted:3:1"
  "payment.payment_order_created:3:1"
  "payment.payment_order_created.delayed:2:1"
  "shipping.shipping_created:3:1"
  "shipping.shipping_updated:3:1"
)

for t in "${TOPICS[@]}"; do
  IFS=":" read -r name partitions rf <<< "$t"
  kafka-topics --create \
    --if-not-exists \
    --bootstrap-server $KAFKA_BROKER \
    --replication-factor $rf \
    --partitions $partitions \
    --topic $name
done

echo ">>> Done creating topics."
