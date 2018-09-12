#!/bin/bash

# Regenerate compiled go protos.
#
# Requires protoc and protoc-gen-go to be installed.
# For more info on protoc:
#    https://developers.google.com/protocol-buffers/
#
# To install protoc-gen-go:
#    go get -u github.com/golang/protobuf/protoc-gen-go

set -e

protoc --go_out=paths=source_relative:. transformer/request/request.proto
