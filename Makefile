# Load application environment variables
include ./deployments/.env
export $(shell sed 's/=.*//' deployments/.env)

# Create Network for application
network:
	docker network create -d bridge ${NETWORK_NAME} || true

# Create and pull the required docker images
docker: network
	docker pull golang:${GO_VERSION}-alpine${ALPINE_VERSION}
	docker pull postgres:alpine${ALPINE_VERSION}
	docker-compose -f ./deployments/docker-compose.yml build app

# Build the application
build:
	go mod download
	go mod verify
	go mod tidy -v
	go build -o ${APP_NAME} ./cmd/http-rest-api/...

# Run Go application container 
app: network
	docker-compose -f ./deployments/docker-compose.yml up -d --build app

# Stop Go application
stop:
	docker compose -f ./deployments/docker-compose.yml stop app

# Remove Go application container
remove: stop
	docker compose -f ./deployments/docker-compose.yml rm --force app

# Remove Docker containers and generated images
clean: stop remove
	docker compose -f ./deployments/docker-compose.yml stop db
	docker compose -f ./deployments/docker-compose.yml rm --force db
	docker rmi -f ${APP_NAME}:${GO_VERSION}-${ALPINE_VERSION} || true
	docker network rm ${NETWORK_NAME} || true
	docker images -f dangling=true -q | xargs docker rmi

# TODO: Implement database handling

# Generate mock objects
mocks:
	mockery --name=UserRepository --srcpkg=./internal/domain/repository --output=./internal/domain/repository/mocks
	mockery --name=UserService --srcpkg=./internal/domain/service --output=./internal/domain/service/mocks
	mockery --name=PgConn --srcpkg=./internal/infrastructure/repository --output=./internal/infrastructure/repository/mocks
