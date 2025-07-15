# !/bin/sh
set -xe

go test ./... -cover -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html

if [ "$1" = "html" ]; then
  explorer cover.html
fi
