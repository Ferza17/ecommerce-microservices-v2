#!/bin/sh

set -e

echo "Running Migration UP Command"
./product-service migration up
echo "Running Migration Elasticsearch Command"
./product-service migration elasticsearch
echo "Running product service"
./product-service run