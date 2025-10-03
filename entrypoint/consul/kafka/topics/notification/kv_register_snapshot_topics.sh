#!/bin/sh

echo "INIT CONFIG KAFKA NAMESPACE NOTIFICATION SNAPSHOT TOPICS"

## LOCAL
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_USER_LOGIN "snapshot-notifications-email_otp_user_login"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_OTP_USER_LOGIN "confirm-snapshot-notifications-email_otp_user_login"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_OTP_USER_LOGIN "compensate-snapshot-notifications-email_otp_user_login"

consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_USER_REGISTER "snapshot-notifications-email_otp_user_register"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_OTP_USER_REGISTER "confirm-snapshot-notifications-email_otp_user_register"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_OTP_USER_REGISTER "compensate-snapshot-notifications-email_otp_user_register"

consul kv put local/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "snapshot-notifications-email_payment_order_created"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_PAYMENT_ORDER_CREATED "confirm-snapshot-notifications-email_payment_order_created"
consul kv put local/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_PAYMENT_ORDER_CREATED "compensate-snapshot-notifications-email_payment_order_created"

## PRODUCTION
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_USER_LOGIN "snapshot-notifications-email_otp_user_login"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_OTP_USER_LOGIN "confirm-snapshot-notifications-email_otp_user_login"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_OTP_USER_LOGIN "compensate-snapshot-notifications-email_otp_user_login"

consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_OTP_USER_REGISTER "snapshot-notifications-email_otp_user_register"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_OTP_USER_REGISTER "confirm-snapshot-notifications-email_otp_user_register"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_OTP_USER_REGISTER "compensate-snapshot-notifications-email_otp_user_register"

consul kv put production/broker/kafka/TOPICS/NOTIFICATION/EMAIL_PAYMENT_ORDER_CREATED "snapshot-notifications-email_payment_order_created"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/CONFIRM/EMAIL_PAYMENT_ORDER_CREATED "confirm-snapshot-notifications-email_payment_order_created"
consul kv put production/broker/kafka/TOPICS/NOTIFICATION/COMPENSATE/EMAIL_PAYMENT_ORDER_CREATED "compensate-snapshot-notifications-email_payment_order_created"
echo "DONE INIT CONFIG KAFKA NAMESPACE NOTIFICATION SNAPSHOT TOPICS"
