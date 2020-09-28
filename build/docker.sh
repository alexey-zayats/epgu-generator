#!/usr/bin/env bash

REGISTRY_URL=aazayats
IMAGE=${REGISTRY_URL}/epgu-generator
VERSION=$(cat VERSION)

docker build -t ${IMAGE} .
docker tag ${IMAGE}:latest ${IMAGE}:${VERSION}

docker push ${IMAGE}:${VERSION}
docker push ${IMAGE}:latest