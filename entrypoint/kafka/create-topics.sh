#!/bin/bash

echo ">>> Creating Kafka topics---"
KAFKA_BROKER="kafka-local-broker-1:29092"

SNAPSHOT_TOPICS=(
  "snapshot-users-user_created:3:1"
  "confirm_snapshot-users-user_created:3:1"
  "compensate_snapshot-users-user_created:3:1"

  "snapshot-users-user_updated:3:1"
  "confirm_snapshot-users-user_updated:3:1"
  "compensate_snapshot-users-user_updated:3:1"

  "snapshot-users-user_login:3:1"
  "confirm_snapshot-users-user_login:3:1"
  "compensate_snapshot-users-user_login:3:1"

  "snapshot-users-user_logout:3:1"
  "confirm_snapshot-users-user_logout:3:1"
  "compensate_snapshot-users-user_logout:3:1"

  "snapshot-products-product_created:3:1"
  "confirm_snapshot-products-product_created:3:1"
  "compensate_snapshot-products-product_created:3:1"

  "snapshot-products-product_updated:3:1"
  "confirm_snapshot-products-product_updated:3:1"
  "compensate_snapshot-products-product_updated:3:1"

  "snapshot-products-product_deleted:3:1"
  "confirm_snapshot-products-product_deleted:3:1"
  "compensate_snapshot-products-product_deleted:3:1"

  "snapshot-notifications-email_otp_created:2:1"
  "confirm_snapshot-notifications-email_otp_created:2:1"
  "compensate_snapshot-notifications-email_otp_created:2:1"

  "snapshot-notifications-email_payment_order_created:2:1"
  "confirm_snapshot-notifications-email_payment_order_created:2:1"
  "compensate_snapshot-notifications-email_payment_order_created:2:1"

  "snapshot-commerce-cart_created:3:1"
  "confirm_snapshot-commerce-cart_created:3:1"
  "compensate_snapshot-commerce-cart_created:3:1"

  "snapshot-commerce-cart_updated:3:1"
  "confirm_snapshot-commerce-cart_updated:3:1"
  "compensate_snapshot-commerce-cart_updated:3:1"

  "snapshot-commerce-cart_deleted:3:1"
  "confirm_snapshot-commerce-cart_deleted:3:1"
  "compensate_snapshot-commerce-cart_deleted:3:1"

  "snapshot-commerce-wishlist_created:3:1"
  "confirm_snapshot-commerce-wishlist_created:3:1"
  "compensate_snapshot-commerce-wishlist_created:3:1"

  "snapshot-commerce-wishlist_updated:3:1"
  "confirm_snapshot-commerce-wishlist_updated:3:1"
  "compensate_snapshot-commerce-wishlist_updated:3:1"

  "snapshot-commerce-wishlist_deleted:3:1"
  "confirm_snapshot-commerce-wishlist_deleted:3:1"
  "compensate_snapshot-commerce-wishlist_deleted:3:1"

  "snapshot-payments-payment_order_created:3:1"
  "confirm_snapshot-payments-payment_order_created:3:1"
  "compensate_snapshot-payments-payment_order_created:3:1"

  "snapshot-payments-payment_order_cancelled_delayed:2:1"
  "confirm_snapshot-payments-payment_order_cancelled_delayed:2:1"
  "compensate_snapshot-payments-payment_order_cancelled_delayed:2:1"

  "snapshot-shippings-shipping_created:3:1"
  "confirm_snapshot-shippings-shipping_created:3:1"
  "compensate_snapshot-shippings-shipping_created:3:1"

  "snapshot-shippings-shipping_updated:3:1"
  "confirm_snapshot-shippings-shipping_updated:3:1"
  "compensate_snapshot-shippings-shipping_updated:3:1"

)

for t in "${SNAPSHOT_TOPICS[@]}"; do
  IFS=":" read -r name partitions rf <<< "$t"
  kafka-topics --create \
    --if-not-exists \
    --bootstrap-server $KAFKA_BROKER \
    --replication-factor $rf \
    --partitions $partitions \
    --topic $name
done

echo ">>> Done creating topics-"
