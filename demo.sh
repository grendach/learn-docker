# Show docker networks
docker network ls
# NETWORK ID          NAME                DRIVER
# 7fca4eb8c647        bridge              bridge
# 9f904ee27bf5        none                null
# cf03ee007fb4        host                host

# Build docker image
docker build -t demo:demo .

# Run docker in bridge network mode
docker run  -p 8080:8000 -it demo:demo
docker run  -p 8090:9000 -it demo:demo

#Run docker in none network mode
docker run  --network=none -p 8090:8000 demo:demo

#Run docker in host network mode
docker run  --network=host  demo:demo

#Check docker network config
docker exec 3b2b5299d17c ip route
docker exec 3b2b5299d17c ip link show