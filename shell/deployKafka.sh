#!/bin/bash

cd ..

docker stack deploy -c kafka.yml kafka

cd shell