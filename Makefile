# Load app environment variables
include ./deployments/app.env
export $(shell sed 's/=.*//' deployments/app.env)

# Create Network for application
network:
	docker network create -d bridge ${NETWORK_NAME} || true

# Build the application
build:
	go mod download
	go mod verify
	go mod tidy -v
	go build -o http-api ./cmd/http-rest-api/...

# Run Go application container 
app: network
	docker-compose --env-file ./deployments/app.env -f ./deployments/docker-compose.yml up -d --build app

# Stop Go application
stop:
	docker compose --env-file ./deployments/app.env -f ./deployments/docker-compose.yml stop app

# Remove Go application container
remove: stop
	docker compose --env-file ./deployments/app.env -f ./deployments/docker-compose.yml rm --force app

# Remove containers and docker images
clean: stop remove
	docker rmi -f ${CONTAINER}:${GO_VERSION}-${DEBIAN_VERSION} || true
	docker network rm ${NETWORK_NAME} || true
	docker images -f dangling=true -q | xargs docker rmi
