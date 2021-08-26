#!/bin/bash

# Go
protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    "$@"


# JS
JSDIR="../ui"
PROTOC_WEB_DIR=$(readlink -f "$JSDIR/node_modules/protoc-gen-grpc-web/bin")
if ! [[ -f "$PROTOC_WEB_DIR/protoc-gen-grpc-web" ]]; then
    echo "$PROTOC_WEB_DIR does not exist, skipping JS build. Did you run 'yarn install' in the ui dir?"
    exit 1
fi

PATH=$PATH:$PROTOC_WEB_DIR
JSPROTO="$JSDIR/src/proto"
mkdir -p $JSPROTO

protoc \
    --js_out=$JSPROTO --js_opt=import_style=commonjs,binary \
    --grpc-web_out=$JSPROTO --grpc-web_opt=import_style=typescript,mode=grpcweb \
    "$@"
