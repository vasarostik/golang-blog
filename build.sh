#!/usr/bin/env bash
go build -o ./package/api/api ./cmd/api
go build -o ./package/migration/migration ./cmd/migration

exec "$@"