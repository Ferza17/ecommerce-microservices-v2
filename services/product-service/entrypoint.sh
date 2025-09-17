#!/bin/sh

set -e

echo "Running Migration UP Command"
./product-service migration up

echo "Running Insert Mock data to Postgresql and Elasticsearch via Sink Connector"
./product-service insert-mock

echo "Running product service"
./product-service run