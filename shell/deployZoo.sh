#!/bin/bash

cd ..

docker stack deploy -c zookeeper.yml zoo

cd shell