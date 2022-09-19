#!/bin/bash
cd ..

docker stack deploy -c <(docker-compose -f traefik.yml config) traefik

cd shell