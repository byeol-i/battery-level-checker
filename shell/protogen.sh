#!/bin/bash
cd ..

protoc -I=. --go_out=plugins=grpc:../../.. pb/svc/auth/*.proto

cd  shell