# Show docker networks
docker network ls

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

# To delete all containers including its volumes use,
docker rm -vf $(docker ps -a -q)
# To delete all the images
docker rmi -f $(docker images -a -q)