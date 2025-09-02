#!/bin/sh

# Makefile
.PHONY: all generate-protos copy-protos
# ==============================================================================
# Main Target
# ==============================================================================


all: yq-eval clean-gen clean-docs generate-protos generate-descriptor copy-protos convert-to-k8s generate-ssl-cert

## PROTO

yq-eval:
	sh ./config/sh/yq-eval.sh

clean-gen:
	sh ./config/sh/clean-gen.sh

clean-docs:
	sh ./config/sh/clean-docs.sh

generate-protos:
	sh ./config/sh/generate-protos.sh

generate-descriptor:
	sh ./config/sh/generate-descriptor.sh

copy-protos:
	sh ./config/sh/copy-protos.sh

## K8S
convert-to-k8s:
	sh ./config/sh/convert-to-k8s.sh

## SSL
generate-ssl-cert:
	sh ./config/sh/generate-ssl-cert.sh