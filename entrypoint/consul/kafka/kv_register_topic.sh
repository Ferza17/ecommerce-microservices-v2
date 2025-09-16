#!/bin/sh

echo "INIT CONFIG KAFKA SNAPSHOT TOPICS"
## user local
consul kv put local/broker/kafka/TOPICS/USER/USER_CREATED "snapshot-users-user_created"
consul kv put local/broker/kafka/TOPICS/USER/USER_UPDATED "snapshot-users-user_updated"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGIN "snapshot-users-user_login"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGOUT "snapshot-users-user_logout"

## user production
consul kv put production/broker/kafka/TOPICS/USER/USER_CREATED "snapshot-users-user_created"
consul kv put production/broker/kafka/TOPICS/USER/USER_UPDATED "snapshot-users-user_updated"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGIN "snapshot-users-user_login"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGOUT "snapshot-users-user_logout"

## product local
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "snapshot-products-product_created"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "snapshot-products-product_updated"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "snapshot-products-product_deleted"

## product production
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "snapshot-products-product_created"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "snapshot-products-product_updated"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "snapshot-products-product_deleted"

## notification local
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "snapshot-notifications-email_otp_created"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "snapshot-notifications-email_payment_order_created"

## notification production
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "snapshot-notifications-email_otp_created"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "snapshot-notifications-email_payment_order_created"

## commerce local
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_CREATED "snapshot-commerce-cart_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "snapshot-commerce-cart_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"
consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_CREATED "snapshot-commerce-wishlist_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_UPDATED "snapshot-commerce-wishlist_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"

## commerce production
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_CREATED "snapshot-commerce-cart_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "snapshot-commerce-cart_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_DELETED "snapshot-commerce-cart_deleted"
consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_CREATED "snapshot-commerce-wishlist_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_UPDATED "snapshot-commerce-wishlist_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/WISHLIST_DELETED "snapshot-commerce-wishlist_deleted"

## payment local
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "snapshot-payments-payment_order_created"
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "snapshot-payments-payment_order_cancelled_delayed"

## payment production
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "snapshot-payments-payment_order_created"
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "snapshot-payments-payment_order_cancelled_delayed"

## shipping local
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "snapshot-shippings-shipping_created"
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "snapshot-shippings-shipping_updated"

## shipping production
consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "snapshot-shippings-shipping_created"
consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "snapshot-shippings-shipping_updated"

echo "DONE INIT CONFIG KAFKA SNAPSHOT TOPICS"
