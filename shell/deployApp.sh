#!/bin/bash
cd ..

docker secret create firebase-key ./conf/firebase/key.json

docker stack deploy -c <(docker-compose -f db.yml config) battery
docker stack deploy -c <(docker-compose -f app.yml config) battery

cd shell