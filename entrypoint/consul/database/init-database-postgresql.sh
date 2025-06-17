#!/bin/sh

initialize_database_postgresql(){
  echo "INIT CONFIG DATABASE POSTGRES"

  ## Postgres Local
  curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_USERNAME
  curl --request PUT --data '1234' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_PASSWORD
  curl --request PUT --data 'disable' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_SSL_MODE
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_HOST
  curl --request PUT --data '5432' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_PORT
  curl --request PUT --data 'products' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS
  curl --request PUT --data 'users' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/USERS
  curl --request PUT --data 'payments' http://localhost:8500/v1/kv/local/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS

  ## Postgres Production
  curl --request PUT --data 'postgres' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_USERNAME
  curl --request PUT --data '1234' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_PASSWORD
  curl --request PUT --data 'disable' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_SSL_MODE
  curl --request PUT --data 'postgres-local' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_HOST
  curl --request PUT --data '5432' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_PORT
  curl --request PUT --data 'products' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS
  curl --request PUT --data 'users' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/USERS
  curl --request PUT --data 'payments' http://localhost:8500/v1/kv/production/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS

  echo "DONE INIT CONFIG DATABASE POSTGRES"
}