#!/bin/bash

cd ..

docker build -t apid -f ./cmd/apid/Dockerfile .
docker build -t auth -f ./cmd/auth/Dockerfile .
# docker build -t producer -f ./cmd/producer/Dockerfile .
docker build -t consumer -f ./cmd/consumer/Dockerfile .

#docker build -t primary-db -f ./cmd/db/Dockerfile.primary .
#docker build -t replica-db -f ./cmd/db/Dockerfile.replica .

cd shell