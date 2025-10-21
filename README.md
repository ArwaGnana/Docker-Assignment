# Run locally

```
go mod init web-app
go mod tidy
go build -o web-app
go run main.go
```

# Using Docker

## Create Docker Image

docker build -t gandalf-colombo-app:v1.0 .

docker images | grep gandalf

## Run Container

docker run -d -p 8000:8080 --name gandalf-app gandalf-colombo-app:v1.0

docker ps

## Test application

localhost:8000

## Inside the container

docker exec -it gandalf-app sh

# Push to DockerHub

docker login

docker tag gandalf-colombo-app:v1.0 arwagnana/gandalf-colombo-app:v1.0

docker tag gandalf-colombo-app:v1.0 arwagnana/gandalf-colombo-app:latest

docker push arwagnana/gandalf-colombo-app:v1.0
docker push arwagnana/gandalf-colombo-app:latest

# Run Container from image on Dockerhub

docker pull arwagnana/gandalf-colombo-app:latest

docker run -d -p 8080:8080 --name gandalf-app arwagnana/gandalf-colombo-app:latest

# Remove everything

## Stop and remove container

docker stop gandalf-app
docker rm gandalf-app

## Remove all local images

docker rmi gandalf-colombo-app:v1.0
docker rmi arwagnana/gandalf-colombo-app:v1.0
docker rmi arwagnana/gandalf-colombo-app:latest

docker images
