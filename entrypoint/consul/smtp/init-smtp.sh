#!/bin/sh

initialize_smtp(){
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

  echo "DONE INIT CONFIG SMTP"
}