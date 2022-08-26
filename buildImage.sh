#!/bin/bash

docker build -t apid -f ./cmd/apid/Dockerfile .
docker build -t consumer -f ./cmd/consumer/Dockerfile .
docker build -t primary-dB -f ./cmd/db/Dockerfile.primary .
docker build -t replica-dB -f ./cmd/db/Dockerfile.replica .
