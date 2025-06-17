#!/bin/sh

initialize_database_mongodb(){
  echo "INIT CONFIG DATABASE MONGODB"

  ### MongoDB Local
  curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_USERNAME
  curl --request PUT --data '1234' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_PASSWORD
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_HOST
  curl --request PUT --data '27017' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_PORT
  curl --request PUT --data 'event-store' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE
  curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/COMMERCE
  curl --request PUT --data 'notification' http://localhost:8500/v1/kv/local/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION
  ## MongoDB Production
  curl --request PUT --data 'mongo' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_USERNAME
  curl --request PUT --data '1234' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_PASSWORD
  curl --request PUT --data 'mongo-local' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_HOST
  curl --request PUT --data '27017' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_PORT
  curl --request PUT --data 'event-store' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE
  curl --request PUT --data 'commerce' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/COMMERCE
  curl --request PUT --data 'notification' http://localhost:8500/v1/kv/production/database/mongodb/MONGO_DATABASE_NAME/NOTIFICATION

  echo "DONE CONFIG DATABASE MONGODB"

}