#!/bin/bash

echo ">>> Creating Kafka topics..."
KAFKA_BROKER="kafka-local-broker-1:9092"

# Define topics (name:partitions:replication-factor)
TOPICS=(
  "sink.pg.users.users:3:1"
  "sink.pg.users.roles:3:1"
  "users.user_created.snapshot:3:1"
  "users.user_updated.snapshot:3:1"
  "users.user_login.snapshot:3:1"
  "users.user_logout.snapshot:3:1"

  "sink.pg.products.products:3:1"
  "sink.es.products.products:3:1"
  "products.product_created.snapshot:3:1"
  "products.product_updated.snapshot:3:1"
  "products.product_deleted.snapshot:3:1"

  "source.mongo.notifications.notification_templates:3:1"
  "notifications.email_otp_created.snapshot:2:1"
  "notifications.email_payment_order_created.snapshot:2:1"

  "commerce.cart_created.snapshot:3:1"
  "commerce.cart_updated.snapshot:3:1"
  "commerce.cart_deleted.snapshot:3:1"
  "commerce.wishlist_created.snapshot:3:1"
  "commerce.wishlist_updated.snapshot:3:1"
  "commerce.wishlist_deleted.snapshot:3:1"

  "sink.pg.payments.payments:3:1"
  "payments.payment_order_created.snapshot:3:1"
  "payments.payment_order_cancelled_delayed.snapshot:2:1"

  "sink.pg.shippings.shippings:3:1"
  "sink.pg.shippings.shipping_providers:3:1"
  "shippings.shipping_created.snapshot:3:1"
  "shippings.shipping_updated.snapshot:3:1"
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
