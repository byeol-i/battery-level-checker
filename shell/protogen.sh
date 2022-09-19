#!/bin/bash
cd ..

protoc -I=. --go_out=plugins=grpc:../../.. pb/svc/firebase/*.proto

cd  shell