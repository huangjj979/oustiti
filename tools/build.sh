#!/usr/bin/env bash

go build -o artifacts/outstiti .
docker build -t outstiti -f tools/Dockerfile .
docker save -o artifacts/outstiti_docker.tar outstiti