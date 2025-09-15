#!/bin/sh
echo "INIT KSQLDB init_user_stream"
ksql http://ksqldb-server-local:8088 < /streams/user/init_user_stream.sql