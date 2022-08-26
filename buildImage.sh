#!/bin/bash

docker build -t apid -f ./cmd/apid/Dockerfile .
docker build -t consumer -f ./cmd/consumer/Dockerfile .
docker build -t primaryDB -f ./cmd/db/Dockerfile.primary .
docker build -t replicaDB -f ./cmd/db/Dockerfile.replica .
