#!/bin/bash

docker stack deploy -c <(docker-compose -f traefik.yml config) traefik
docker stack deploy -c <(docker-compose -f app.yml config) app
#docker stack deploy -c traefik.yml traefik
