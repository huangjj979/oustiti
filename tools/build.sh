#!/usr/bin/env bash

set -e

export CGO_ENABLED=0

go build -tags netgo -v -o artifacts/outstiti .
#docker build -t outstiti -f tools/Dockerfile .
#docker save -o artifacts/outstiti_docker.tar outstiti