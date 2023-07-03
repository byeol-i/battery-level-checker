#!/bin/bash
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

docker stack deploy -c <(docker-compose -f db.yml config) battery
docker stack deploy -c <(docker-compose -f app.yml config) battery

cd shell