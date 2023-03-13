#!/bin/bash
cd ..

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