#!/usr/bin/env bash
go build -o ./dockerfiles/api/api ./cmd/api
go build -o ./dockerfiles/migration/migration ./cmd/migration
go build -o ./dockerfiles/grpc/grpc ./cmd/grpc
go build -o ./dockerfiles/nats/nats ./cmd/nats

exec "$@"