#!/bin/sh


echo "=================================================="
echo ">>> Generate Certificate with mkcert <<<"
echo "=================================================="
mkdir -p ./entrypoint/traefik/certs \

# Generate wildcard certificate for your local domains
mkcert -cert-file ./entrypoint/traefik/certs/localhost-cert.pem \
           -key-file ./entrypoint/traefik/certs/localhost-key.pem \
           "*.local.dev" "local.dev" "*.localhost" "localhost" "127.0.0.1" \

mkcert --install

echo "=================================================="
echo ">>> Done. Generate Certificate with mkcert <<<"
echo "=================================================="