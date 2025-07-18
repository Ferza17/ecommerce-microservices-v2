#!/bin/sh

echo "INIT CONSUL KV DATABASE MONGODB"

### MongoDB Local
consul kv put local/database/mongodb/MONGO_USERNAME "mongo"
consul kv put local/database/mongodb/MONGO_PASSWORD "1234"
consul kv put local/database/mongodb/MONGO_HOST "localhost"
consul kv put local/database/mongodb/MONGO_PORT "27017"
## DB NAME
consul kv put local/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE "event-store"
consul kv put local/database/mongodb/MONGO_DATABASE_NAME/COMMERCE "commerce"
consul kv put local/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION "notification"

## MongoDB Production
consul kv put production/database/mongodb/MONGO_USERNAME "mongo"
consul kv put production/database/mongodb/MONGO_PASSWORD "1234"
consul kv put production/database/mongodb/MONGO_HOST "mongo-local"
consul kv put production/database/mongodb/MONGO_PORT "27017"
## DB NAME
consul kv put production/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE "event-store"
consul kv put production/database/mongodb/MONGO_DATABASE_NAME/COMMERCE "commerce"
consul kv put production/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION "notification"

echo "✅ INIT CONSUL KV DATABASE MONGODB"
