

echo "INIT CONSUL KV DATABASE POSTGRES"

## Postgres Local
consul kv put local/database/postgres/POSTGRES_USERNAME "postgres"
consul kv put local/database/postgres/POSTGRES_PASSWORD "1234"
consul kv put local/database/postgres/POSTGRES_SSL_MODE "disable"
consul kv put local/database/postgres/POSTGRES_HOST "localhost"
consul kv put local/database/postgres/POSTGRES_PORT "5432"

# Database Name
consul kv put local/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS "products"
consul kv put local/database/postgres/POSTGRES_DATABASE_NAME/USERS "users"
consul kv put local/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS "payments"

## Postgres Production
consul kv put production/database/postgres/POSTGRES_USERNAME "postgres"
consul kv put production/database/postgres/POSTGRES_PASSWORD "1234"
consul kv put production/database/postgres/POSTGRES_SSL_MODE "disable"
consul kv put production/database/postgres/POSTGRES_HOST "postgres-local"
consul kv put production/database/postgres/POSTGRES_PORT "5432"

# Database Name
consul kv put production/database/postgres/POSTGRES_DATABASE_NAME/PRODUCTS "products"
consul kv put production/database/postgres/POSTGRES_DATABASE_NAME/USERS "users"
consul kv put production/database/postgres/POSTGRES_DATABASE_NAME/PAYMENTS "payments"

echo "✅ DONE INIT CONSUL KV DATABASE POSTGRES"
