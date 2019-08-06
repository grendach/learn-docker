# Show docker networks
docker network ls
# NETWORK ID          NAME                DRIVER
# 7fca4eb8c647        bridge              bridge
# 9f904ee27bf5        none                null
# cf03ee007fb4        host                host

docker build -t docker-demo:1111 .
# docker build -t docker-demo:2222 .
# docker build -t docker-demo:3333 .

#Запуск контейнера в режимі Мосту:
docker run --network=bridge -p 8080:5000 -it docker:demo
docker run --network=bridge -p 8090:5000 -it docker:demo

#Запуск в режимі Ніякому
docker run --network=none -p 2222:3001 docker-demo:1111

#Запуск в режимі хоста
docker run --network=host  docker-demo:3333