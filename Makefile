#!/bin/sh

# Makefile
.PHONY: all generate-protos copy-protos
# ==============================================================================
# Main Target
# ==============================================================================

all: yq-eval clean-gen generate-protos copy-protos convert-to-k8s

# STEP 1, ADD CONFIGURATION GENERATED FILE
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

# STEP 2, CLEAN GENERATED FILE
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

# STEP 3, GENERATED PROTO FILE
generate-protos:
	@echo "=================================================="
	@echo ">>> Generating file .proto <<<"
	@echo "=================================================="
	buf generate
	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="

# STEP 3, COPY PROTO FILE TO NODE.JS SERVICES
copy-protos:
	@echo "=================================================="
	@echo ">>> copy file .proto to commerce-service/proto <<<"
	@echo "=================================================="


	# COPY TO COMMERCE SERVICE
	TARGET_DIR="commerce-service/proto"; \
	mkdir -p "$$TARGET_DIR"; \
	rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$$TARGET_DIR"; \

	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="

# STEP 4, GENERATED DOCKER COMPOSE TO K8S MANIFEST
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
