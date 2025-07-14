#!/bin/sh


echo "=================================================="
echo ">>> copy file .proto to commerce-service/proto <<<"
echo "=================================================="

# COPY TO COMMERCE SERVICE
TARGET_DIR="commerce-service/proto";
mkdir -p "$TARGET_DIR";
rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$TARGET_DIR";

echo "=================================================="
echo ">>> copy file .proto to shipping-service/proto <<<"
echo "=================================================="

TARGET_DIR="shipping-service/proto";
mkdir -p "$TARGET_DIR";
rsync -av --include='*.proto' --include='*/' --exclude='*' "proto/" "$TARGET_DIR";

echo "=================================================="
echo ">>> Done. <<<"
echo "=================================================="