#!/usr/bin/env bash

PROTO_NAMES=(
  "im"
)

for name in "${PROTO_NAMES[@]}"; do
  protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      ${name}/${name}.proto
  if [ $? -ne 0 ]; then
      echo "error processing ${name}.proto"
      exit $?
  fi
done
