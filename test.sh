# !/bin/sh
set -xe

go test ./... -cover -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html

if [ $# -eq 1 ]; then
  explorer cover.html
fi
