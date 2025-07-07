#!/bin/sh

# Makefile
.PHONY: all generate-protos copy-protos
# ==============================================================================
# Main Target
# ==============================================================================

all: yq-eval clean-gen clean-docs generate-protos generate-descriptor copy-protos convert-to-k8s generate-ssl-cert


yq-eval:
	@echo "Merging config/gen/*.yml into buf.gen.yaml"

	mkdir config/buf/plugins/temp_all; \
	yq eval-all '[.]' config/buf/plugins/api-gateway/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/api_gateway_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/commerce-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/commerce_service_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/event-store-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/event_store_service_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/notification-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/notification_service_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/payment-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/payment_service_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/product-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/product_service_plugins.yml; \
	yq eval-all '[.]' config/buf/plugins/user-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/user_service_plugins.yml; \

	yq eval-all '[.]' config/buf/plugins/temp_all/*.yml | yq eval 'flatten' > config/buf/plugins/all_plugins.yml; \
	yq eval-all '.plugins = load("config/buf/plugins/all_plugins.yml") | .' config/buf/_base_config.yml > buf.gen.yaml; \
	rm -r config/buf/plugins/all_plugins.yml;\
	rm -r config/buf/plugins/temp_all;\


	@echo "Done."

clean-gen:
	@echo "Cleaning proto generated file"

	rm -r api-gateway/model/rpc/gen; \
	rm -r commerce-service/src/model/rpc/gen; \
	rm -r event-store-service/model/rpc/gen; \
	rm -r notification-service/model/rpc/gen; \
	rm -r payment-service/model/rpc/gen; \
	rm -r product-service/model/rpc/gen; \
	rm -r user-service/model/rpc/gen; \

	@echo "Done Cleaning proto generated file"

clean-docs:
	@echo "Cleaning docs generated file"
	rm -r product-service/docs; \
	rm -r user-service/docs; \

	@echo "Done Cleaning docs generated file"

generate-protos:
	@echo "=================================================="
	@echo ">>> Generating file .proto <<<"
	@echo "=================================================="
	buf generate
	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="

generate-descriptor:
	@echo "Generating descriptor.pb"

	buf build --as-file-descriptor-set --output api-gateway/descriptor.pb;\
	buf build --as-file-descriptor-set --output commerce-service/descriptor.pb;\
	buf build --as-file-descriptor-set --output event-store-service/descriptor.pb;\
	buf build --as-file-descriptor-set --output notification-service/descriptor.pb;\
	buf build --as-file-descriptor-set --output payment-service/descriptor.pb;\
	buf build --as-file-descriptor-set --output product-service/descriptor.pb;\
	buf build --as-file-descriptor-set --output user-service/descriptor.pb;\

	@echo "Done Generating descriptor.pb"

copy-protos:
	@echo "=================================================="
	@echo ">>> copy file .proto to commerce-service/proto <<<"
	@echo "=================================================="


	# COPY TO COMMERCE SERVICE
	TARGET_DIR_COMMERCE_SERVICE="commerce-service/proto"; \
	mkdir -p "$$TARGET_DIR_COMMERCE_SERVICE"; \
	rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$$TARGET_DIR_COMMERCE_SERVICE"; \

	# COPY PROTO TO SHIPPING SERVICE
	TARGET_DIR_SHIPPING_SERVICE="shipping-service/src/main/proto"; \
	mkdir -p "$$TARGET_DIR_SHIPPING_SERVICE"; \
	rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$$TARGET_DIR_SHIPPING_SERVICE"; \


	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="

convert-to-k8s:
	@echo "=================================================="
	@echo ">>> Convert docker-compose.yml to folder k8s <<<"
	@echo "=================================================="

	# Create k8s directory if it doesn't exist
	K8S_DIR="k8s/"; \
	mkdir -p "$$K8S_DIR"; \
	find "$$K8S_DIR" -name "*.yml" -type f -delete; \

	# Convert docker-compose.yml to Kubernetes manifests
	kompose convert --file docker-compose.yml --out "k8s/"; \

	@echo "=================================================="
	@echo ">>> Done. Kubernetes manifests generated in k8s/ <<<"
	@echo "=================================================="

generate-ssl-cert:
	@echo "=================================================="
	@echo ">>> Generate Certificate with mkcert <<<"
	@echo "=================================================="
	mkdir -p ./entrypoint/traefik/certs \

	# Generate wildcard certificate for your local domains
	mkcert -cert-file ./entrypoint/traefik/certs/localhost-cert.pem \
           -key-file ./entrypoint/traefik/certs/localhost-key.pem \
           "*.local.dev" "local.dev" "*.localhost" "localhost" "127.0.0.1" \

	mkcert --install

	@echo "=================================================="
	@echo ">>> Done. Generate Certificate with mkcert <<<"
	@echo "=================================================="