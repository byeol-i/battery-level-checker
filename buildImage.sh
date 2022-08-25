#!/bin/bash

docker build -t apid -f ./cmd/apid/Dockerfile .
docker build -t consumer -f ./cmd/consumer/Dockerfile .
docker build -t producer -f ./cmd/producer/Dockerfile .