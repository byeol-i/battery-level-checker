#!/bin/bash

cd ..

protoc -I=. --go-grpc_out=../../.. --go_out=../../.. pb/**/**/*.proto

cd  shell 