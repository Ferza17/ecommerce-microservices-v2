#!/bin/sh

set -e

echo "Running Migration UP Command"
./payment-service migration up
echo "Running Insert Mock data to Postgresql and Elasticsearch via Sink Connector"
./payment-service insert-mock
echo "Running payment service"
./payment-service run