#!/bin/bash

cd ./networks/single-node
echo "一、./network.sh up"
./network.sh up

echo "二、./network.sh createChannel"
./network.sh createChannel

echo "三、./network.sh deployCC -ccn basic -ccp ../../test-chaincode -ccl go"
./network.sh deployCC -ccn basic -ccp ../../test-chaincode -ccl go

cd ../../mqtt/docker
echo "四、docker-compose -f docker-compose-1node.yml up -d'"
docker-compose -f docker-compose-1node.yml up -d


