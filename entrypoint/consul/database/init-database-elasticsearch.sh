#!/bin/sh

initialize_database_elasticsearch(){
  echo "INIT CONFIG DATABASE ELASTICSEARCH"
  ## Elasticsearch Local
  curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_USERNAME
  curl --request PUT --data '' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_PASSWORD
  curl --request PUT --data 'localhost' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_HOST
  curl --request PUT --data '9200' http://localhost:8500/v1/kv/local/database/elasticsearch/ELASTICSEARCH_PORT
  ## Elasticsearch Production
  curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_USERNAME
  curl --request PUT --data '' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_PASSWORD
  curl --request PUT --data 'elasticsearch-local' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_HOST
  curl --request PUT --data '9200' http://localhost:8500/v1/kv/production/database/elasticsearch/ELASTICSEARCH_PORT

  echo "DONE INIT CONFIG DATABASE ELASTICSEARCH"

}