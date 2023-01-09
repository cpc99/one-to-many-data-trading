#!/bin/bash

cd ./networks/single-node
echo "一、./network.sh down"
./network.sh down

echo "二、./network.sh createChannel"
docker ps -a

cd ../../application_iot_marketplaces
rm -r keystore
rm -r wallet



