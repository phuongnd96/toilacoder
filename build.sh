#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w'
docker build -t long25nt/toilacoder:latest --no-cache .
docker push long25nt/toilacoder:latest