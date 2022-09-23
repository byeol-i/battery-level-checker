#!/bin/bash
cd ..

docker stack deploy -c <(docker-compose -f dashboard.yml config) dashboard

cd shell