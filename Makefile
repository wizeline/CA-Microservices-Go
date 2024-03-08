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

# Remove Docker containers and generated images
clean: stop remove
	docker rmi -f ${CONTAINER}:${GO_VERSION}-${DEBIAN_VERSION} || true
	docker network rm ${NETWORK_NAME} || true
	docker images -f dangling=true -q | xargs docker rmi

# Generate mock objects
mocks:
	mockery --name=UserRepository --srcpkg=./internal/domain/repository --output=./internal/domain/repository/mocks
	mockery --name=UserService --srcpkg=./internal/domain/service --output=./internal/domain/service/mocks
	mockery --name=PgConn --srcpkg=./internal/infrastructure/repository --output=./internal/infrastructure/repository/mocks
