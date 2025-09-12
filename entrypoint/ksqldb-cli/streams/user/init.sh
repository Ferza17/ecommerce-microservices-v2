#!/bin/sh
echo "INIT KSQLDB init_user_stream"
ksql http://ksqldb-server-local:8088 < /scripts/user/init_user_stream.sql