#!/bin/sh

initialize_database_elasticsearch(){
  echo "INIT CONSUL KV DATABASE ELASTICSEARCH"
  ## Elasticsearch Local
  consul kv put local/database/elasticsearch/ELASTICSEARCH_USERNAME ""
  consul kv put local/database/elasticsearch/ELASTICSEARCH_PASSWORD ""
  consul kv put local/database/elasticsearch/ELASTICSEARCH_HOST "localhost"
  consul kv put local/database/elasticsearch/ELASTICSEARCH_PORT "9200"
  
  ## Elasticsearch Production
  consul kv put production/database/elasticsearch/ELASTICSEARCH_USERNAME ""
  consul kv put production/database/elasticsearch/ELASTICSEARCH_PASSWORD ""
  consul kv put production/database/elasticsearch/ELASTICSEARCH_HOST "elasticsearch-local"
  consul kv put production/database/elasticsearch/ELASTICSEARCH_PORT "9200"

  echo "✅ DONE INIT CONSUL KV DATABASE ELASTICSEARCH"
}

initialize_database_elasticsearch