# Makefile
.PHONY: all generate-protos copy-protos
# ==============================================================================
# Main Target
# ==============================================================================

# Target default: Jalankan semua proses yang diperlukan
all: generate-protos copy-protos

# Target untuk menghasilkan kode Go dari file .proto menggunakan buf
generate-protos:
	@echo "=================================================="
	@echo ">>> Generating file .proto <<<"
	@echo "=================================================="
	buf generate
	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="

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
