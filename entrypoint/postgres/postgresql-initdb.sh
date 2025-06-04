#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
  CREATE DATABASE products;
  CREATE DATABASE users;
  CREATE DATABASE payments;
EOSQL
