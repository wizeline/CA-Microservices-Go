# Load application environment variables
include ./deployments/.env
export $(shell sed 's/=.*//' deployments/.env)

help:
	@echo "Available targets:"
	@echo "  pgdb   - Start PostgreSQL container"
	@echo "  pgdb-stop   - Stop PostgreSQL container"
	@echo "  pgadmin    - Start pgAdmin container"
	@echo "  pgadmin-stop    - Stop pgAdmin container"
	@echo "  clean      - Stop and remove containers"

# Create Network for application
network:
	docker network create -d bridge ${NETWORK_NAME} || true

# Create and pull the required docker images
app-img: network
	docker pull golang:${GO_VERSION}-alpine${ALPINE_VERSION}
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

# PostgreSQL database
pgdb:
	docker pull postgres:alpine${ALPINE_VERSION}
	docker-compose -f ./deployments/docker-compose.yml up -d pgdb

pgdb-stop:
	docker-compose -f ./deployments/docker-compose.yml stop pgdb

pgdb-rm: pgdb-stop
	docker-compose -f ./deployments/docker-compose.yml rm --force pgdb

pgdb-clean: pgdb-rm
	docker rmi -f postgres:alpine${ALPINE_VERSION} || true

# PgAdmin: default PostgreSQL GUI
pgadmin:
	docker pull dpage/pgadmin4:${PGADMIN_VERSION}
	docker-compose -f ./deployments/docker-compose.yml up -d pgadmin

pgadmin-stop:
	docker compose -f ./deployments/docker-compose.yml stop pgadmin

pgadmin-rm: pgadmin-stop
	docker compose -f ./deployments/docker-compose.yml rm --force pgadmin

pgadmin-clean: pgadmin-rm
	docker rmi -f dpage/pgadmin4:${PGADMIN_VERSION} || true

# Generate mock objects
mocks:
	mockery --name=UserRepository --srcpkg=./internal/domain/repository --output=./internal/domain/repository/mocks
	mockery --name=UserService --srcpkg=./internal/domain/service --output=./internal/domain/service/mocks
