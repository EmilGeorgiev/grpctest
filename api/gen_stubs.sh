#!/bin/bash

command_exists () {
    type "$1" &> /dev/null ;
}

if !(command_exists protoc) ; then
	echo "protobuf is required for the stub generation. More information could be taken from: https://developers.google.com/protocol-buffers/"
	exit 1
fi

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc -I inventory/v1/ inventory/v1/*.proto --go_out=plugins=grpc:inventory/v1

echo 'All stubs were generated successfully.'



