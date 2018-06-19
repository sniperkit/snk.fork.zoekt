#!/bin/sh
set -x
set -e

env GOOS=linux go build -v ./docker/zoekt-mirror-gitiles
docker build --force-rm -t sniperkit/zoekt-mirror-gitiles:3.7-alpine --no-cache -f ./docker/dockerfile-alpine3.7