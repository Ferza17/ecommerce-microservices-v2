#!/bin/sh

echo "Registering OPEN POLICY AGENT as service in Consul ..."

# Wait for Mailhog SMTP to be available
echo "Waiting for OPEN POLICY AGENT to be available..."
until nc -z opa-local 9191; do
echo "OPEN POLICY AGENT not ready yet, waiting..."
sleep 2
done
echo "OPEN POLICY AGENT is available"


# Register Mailhog SMTP service
consul services register \
-name=opa \
-id=opa-main \
-port=9191 \
-address=opa-local \
-tag=policy \
-tag=rbac \
-tag=mailhog \


## TODO: HEALTH CHECK