#!/bin/sh

echo "INIT CONFIG OPEN POLICY AGENT"
# LOCAL
consul kv put local/policy/opa/PATH "http://localhost:8181/v1/data/authz/allow"

# PRODUCTION
consul kv put production/policy/opa/PATH "http://opa-local:8181/v1/data/authz/allow"
echo "DONE INIT CONFIG OPEN POLICY AGENT"


