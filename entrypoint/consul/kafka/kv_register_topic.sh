#!/bin/sh

echo "INIT CONFIG KAFKA TOPICS"
## user local
consul kv put local/broker/kafka/TOPICS/USER/SINK_PG "user.users.sink.pg"
consul kv put local/broker/kafka/TOPICS/USER/USER_CREATED "user.user_created.snapshot"
consul kv put local/broker/kafka/TOPICS/USER/USER_UPDATED "user.user_updated.snapshot"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGIN "user.user_login.snapshot"
consul kv put local/broker/kafka/TOPICS/USER/USER_LOGOUT "user.user_logout.snapshot"

## user production
consul kv put production/broker/kafka/TOPICS/USER/SINK_PG "user.users.sink.pg"
consul kv put production/broker/kafka/TOPICS/USER/USER_CREATED "user.user_created.snapshot"
consul kv put production/broker/kafka/TOPICS/USER/USER_UPDATED "user.user_updated.snapshot"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGIN "user.user_login.snapshot"
consul kv put production/broker/kafka/TOPICS/USER/USER_LOGOUT "user.user_logout.snapshot"

## product local
consul kv put local/broker/kafka/TOPICS/PRODUCT/SINK_PG "product.products.sink.pg"
consul kv put local/broker/kafka/TOPICS/PRODUCT/SINK_ES "product.products.sink.es"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "product.product_created.snapshot"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "product.product_updated.snapshot"
consul kv put local/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "product.product_deleted.snapshot"


## product production
consul kv put production/broker/kafka/TOPICS/PRODUCT/SINK_PG "product.products.sink.pg"
consul kv put production/broker/kafka/TOPICS/PRODUCT/SINK_ES "product.products.sink.es"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_CREATED "product.product_created.snapshot"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_UPDATED "product.product_updated.snapshot"
consul kv put production/broker/kafka/TOPICS/PRODUCT/PRODUCT_DELETED "product.product_deleted.snapshot"

## notification local
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/SINK_MONGO "notification.notification_templates.sink.mongo"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "notification.email_otp_created.snapshot"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "notification.email_payment_order_created.snapshot"

## notification production
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/SINK_MONGO "notification.notification_templates.sink.mongo"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_CREATED "notification.email_otp_created.snapshot"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "notification.email_payment_order_created.snapshot"

## commerce local
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_CREATED "commerce.cart_created.snapshot"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "commerce.cart_updated.snapshot"
consul kv put local/broker/kafka/TOPICS/COMMERCE/CART_DELETED "commerce.cart_deleted.snapshot"

## commerce production
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_CREATED "commerce.cart_created.snapshot"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_UPDATED "commerce.cart_updated.snapshot"
consul kv put production/broker/kafka/TOPICS/COMMERCE/CART_DELETED "commerce.cart_deleted.snapshot"

## payment local
consul kv put local/broker/kafka/TOPICS/PAYMENT/SINK_PG "payment.payments.sink.pg"
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "payment.payment_order_created.snapshot"
consul kv put local/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "payment.payment_order_cancelled_delayed.snapshot"

## payment production
consul kv put production/broker/kafka/TOPICS/PAYMENT/SINK_PG "payment.payments.sink.pg"
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED "payment.payment_order_created.snapshot"
consul kv put production/broker/kafka/TOPICS/PAYMENT/PAYMENT_ORDER_CREATED_DELAYED "payment.payment_order_cancelled_delayed.snapshot"

## shipping local
consul kv put local/broker/kafka/TOPICS/SHIPPING/SINK_PG "payment.payments.sink.pg"
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "shipping.shipping_created.snapshot"
consul kv put local/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "shipping.shipping_updated.snapshot"

## shipping production
consul kv put production/broker/kafka/TOPICS/SHIPPING/SINK_PG "payment.payments.sink.pg"
consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_CREATED "shipping.shipping_created.snapshot"
consul kv put production/broker/kafka/TOPICS/SHIPPING/SHIPPING_UPDATED "shipping.shipping_updated.snapshot"
echo "DONE INIT CONFIG KAFKA TOPICS"