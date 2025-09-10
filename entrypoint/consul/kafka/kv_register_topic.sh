#!/bin/sh

echo "INIT CONFIG KAFKA TOPICS"
## user local
consul kv put local/broker/kafka/TOPICS/USER/USER_CREATED "user.user_created"
consul kv put local/broker/kafka/TOPICS/USER/USER_UPDATED "user.user_updated"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGIN "user.user_login"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGOUT "user.user_logout"

## user production
consul kv put production/broker/kafka/TOPICS/USER/USER_CREATED "user.user_created"
consul kv put production/broker/kafka/TOPICS/USER/USER_UPDATED "user.user_updated"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGIN "user.user_login"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGOUT "user.user_logout"

## product local
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "product.product_created"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "product.product_updated"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "product.product_deleted"


## product production
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "product.product_created"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "product.product_updated"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "product.product_deleted"

## notification local
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "notification.email_otp_created"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "notification.email_payment_order_created"

## notification production
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "notification.email_otp_created"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "notification.email_payment_order_created"

## commerce local
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_CREATED "commerce.cart_created"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "commerce.cart_updated"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_DELETED "commerce.cart_deleted"

## commerce production
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_CREATED "commerce.cart_created"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "commerce.cart_updated"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_DELETED "commerce.cart_deleted"


## payment local
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "payment.payment_order_created"
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "payment.payment_order_created.delayed"

## payment production
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "payment.payment_order_created"
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "payment.payment_order_created.delayed"

## commerce local
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "shipping.shipping_created"
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "shipping.shipping_updated"

echo "DONE INIT CONFIG KAFKA TOPICS"