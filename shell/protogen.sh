#!/bin/bash
cd ..

protoc -I=. --go-grpc_out=../../.. --go_out=../../.. pb/svc/firebase/*.proto
cd  shell 