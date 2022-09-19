#!/bin/bash
cd ..

docker stack deploy -c <(docker-compose -f app.yml config) app

cd shell