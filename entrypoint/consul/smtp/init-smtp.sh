#!/bin/sh

echo "INIT CONFIG SMTP"

## Local
consul kv put local/smtp/SMTP_SENDER_EMAIL "ecommerce@email.com"
consul kv put local/smtp/SMTP_USERNAME
consul kv put local/smtp/SMTP_PASSWORD
consul kv put local/smtp/SMTP_HOST "localhost"
consul kv put local/smtp/SMTP_PORT "1025"

## Production
consul kv put production/smtp/SMTP_SENDER_EMAIL "ecommerce@email.com"
consul kv put production/smtp/SMTP_USERNAME
consul kv put production/smtp/SMTP_PASSWORD
consul kv put production/smtp/SMTP_HOST "mailhog-local"
consul kv put production/smtp/SMTP_PORT "1025"

echo "DONE INIT CONFIG SMTP"
