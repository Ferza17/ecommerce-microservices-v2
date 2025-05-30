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

	TARGET_DIR="commerce-service/proto"; \
	mkdir -p "$$TARGET_DIR"; \
	rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$$TARGET_DIR"; \

	@echo "=================================================="
	@echo ">>> Done. <<<"
	@echo "=================================================="
