#!/bin/bash
# build.sh - For Linux builds

set -ex

# Create the target directories
mkdir -p $PREFIX/bin
mkdir -p $PREFIX

# Copy the shared library and plugin file
cp ResetScore.so $PREFIX/
cp ResetScore.pplugin $PREFIX/

# Set proper permissions
chmod 755 $PREFIX/ResetScore.so
chmod 644 $PREFIX/ResetScore.pplugin