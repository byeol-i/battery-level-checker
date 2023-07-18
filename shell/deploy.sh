#!/bin/bash

if ! command -v docker &> /dev/null; then
  echo "Docker is not installed..."
  exit 1
fi

if [ -z "$(docker info | grep Swarm)" ]; then
  echo "Not a Swarm..."
  exit 1
fi

if ! docker node ls --format '{{.Hostname}}' | xargs -I {} docker node inspect --format '{{.Spec.Labels}}' {} | grep -q "kafka"; then
  echo "Can't find node with 'kafka' label..."
  exit 1
fi

num_nodes=$(docker node ls | grep -v "ID" | wc -l)
required_nodes=2

if [ "$num_nodes" -lt "$required_nodes" ]; then
  echo "Can't find number of nodes. Required: $required_nodes, Available: $num_nodes. Exiting..."
  exit 1
fi

cd ..

NETWORK_NAME="kafka-network"

EXISTING_NETWORK=$(docker network ls --filter name=$NETWORK_NAME --format '{{.Name}}')

if [ "$EXISTING_NETWORK" == "$NETWORK_NAME" ]; then
    echo "Network '$NETWORK_NAME' already exists."
else
    echo "Network '$NETWORK_NAME' does not exist. Creating network..."

    docker network create -d overlay --attachable $NETWORK_NAME

    echo "Network '$NETWORK_NAME' created successfully."
fi

CliApiVersion=$(docker version -f '{{.Client.APIVersion}}')
echo "Using api version" $CliApiVersion
if [[ "$OSTYPE" == "darwin"* ]]; then
    find .env -type f -exec sed -i '' -e /^CLI_API_VERSION=/s/=.*/=$CliApiVersion/ {} \;
else
    find .env -type f -exec sed -i -e /^CLI_API_VERSION=/s/=.*/=$CliApiVersion/ {} \;
fi

cd shell

sh ./deployZoo.sh

sh ./deployKafka.sh

sh ./deployApp.sh
