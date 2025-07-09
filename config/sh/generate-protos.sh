#!/bin/sh

echo ">>> Generating .proto files with buf <<<"
buf generate
echo ">>> Done."
