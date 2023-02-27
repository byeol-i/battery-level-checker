#!/bin/bash
cd ..

cd cmd/db/

docker-compose up -d 
cd ../..
# docker stack deploy -c <(docker-compose -f db.yml config) app

cd shell