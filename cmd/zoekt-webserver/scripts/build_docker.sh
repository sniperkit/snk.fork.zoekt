#!/bin/sh
set -x
set -e

env GOOS=linux go build -v ./docker/zoekt-webserver
docker build --force-rm -t sniperkit/zoekt-webserver:3.7-alpine --no-cache -f ./docker/dockerfile-alpine3.7