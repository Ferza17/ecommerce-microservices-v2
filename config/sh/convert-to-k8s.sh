#!/bin/sh


echo "=================================================="
echo ">>> Convert docker-compose.yml to folder k8s <<<"
echo "=================================================="

# Create k8s directory if it doesn't exist
DIR="k8s/"; \
mkdir -p "$DIR"; \
find "$DIR" -name "*.yml" -type f -delete; \

# Convert docker-compose.yml to Kubernetes manifests
kompose convert --file docker-compose.yml --out "$DIR"; \

echo "=================================================="
echo ">>> Done. Kubernetes manifests generated in $DIR <<<"
echo "=================================================="