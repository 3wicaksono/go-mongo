#!/bin/sh
NETWORK_NAME=go-net
docker system prune -f
docker network ls|grep ${NETWORK_NAME} > /dev/null || docker network create ${NETWORK_NAME}
docker rm -f mongo-db
docker rm -f go-mongo

if [[ "$(docker images -q mongo:3.6 2> /dev/null)" == "" ]]; then
  echo "docker images pull mongo:3.6"
  docker pull mongo:3.6
fi

if [[ "$(docker images -q go-mongo:latest 2> /dev/null)" != "" ]]; then
  echo "docker delete images go-mongo:latest"
  docker rmi -f go-mongo:latest
fi

docker run  -d --name mongo-db --hostname mongo-db --network ${NETWORK_NAME} -p 27017:27017 -p 28017:28017 -v "$PWD/scripts:/scripts" -t mongo:3.6
docker exec mongo-db sleep 10;
wait
echo done

docker build -t go-mongo:latest .

docker exec mongo-db sh -c "mongo --port 27017 < /scripts/create-user.js";
docker run -d --name=go-mongo --network ${NETWORK_NAME} -p 8787:8787 -t go-mongo:latest sh -c "/opt/go-mongo/go-mongo serve"
docker exec go-mongo sleep 10;
wait
echo done
