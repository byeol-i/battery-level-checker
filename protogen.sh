#!/bin/bash

protoc -I=. --go_out=plugins=grpc:../../.. pb/svc/auth/*.proto

