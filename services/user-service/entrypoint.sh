#!/bin/sh

set -e

./user-service migration up
./user-service acl
./user-service run