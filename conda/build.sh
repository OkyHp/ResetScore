#!/bin/bash
# !build.sh - For Linux builds

set -ex

# Create the target directories
mkdir -p $PREFIX/bin
mkdir -p $PREFIX

# Copy the shared library, plugin and translation files
cp resetscore.so $PREFIX/
cp resetscore.pplugin $PREFIX/
cp resetscore.yml $PREFIX/

# Set proper permissions
chmod 755 $PREFIX/resetscore.so
chmod 644 $PREFIX/resetscore.pplugin
chmod 644 $PREFIX/resetscore.yml