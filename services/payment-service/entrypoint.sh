#!/bin/sh

set -e

echo "Running Migration UP Command"
./payment-service migration up
echo "Running payment service"
./payment-service run