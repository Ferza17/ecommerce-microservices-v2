#!/bin/sh

set -e

./notification-service migration up
./notification-service run