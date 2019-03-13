#!/bin/sh

export PATH=$PATH:$HOME/go/bin/

protoc -I . ./gateway.proto --go_out=plugins=grpc:. --proto_path .
