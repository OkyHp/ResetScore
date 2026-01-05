#!/usr/bin/bash

docker run --rm \
  --platform linux/amd64 \
  --user "$(id -u):$(id -g)" \
  -e BIN_NAME=ResetScore.so \
  -e BUILD_DEBUG=false \
  -v "$(pwd):/app" \
  -w /app \
  steamrt-go-builder