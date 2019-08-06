# Show docker networks
docker network ls
# NETWORK ID          NAME                DRIVER
# 7fca4eb8c647        bridge              bridge
# 9f904ee27bf5        none                null
# cf03ee007fb4        host                host

# Build docker image
docker build -t docker:demo .

# Run docker in bridge network mode
docker run --network=bridge -p 8080:5000 -it docker:demo
docker run --network=bridge -p 8090:5000 -it docker:demo

#Run docker in none network mode
docker run --network=none -p 8090:5000 docker:demo

#Run docker in host network mode
docker run --network=host  docker:demo