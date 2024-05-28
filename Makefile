# Load application environment variables
include ./deployments/.env
export $(shell sed 's/=.*//' deployments/.env)

# Application ---------------------------------------------
.PHONY: network	app-img	build app start stop restart remove clean logs-db

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
start:
	docker-compose -f ./deployments/docker-compose.yml up -d app

# Stop Go application
stop:
	docker compose -f ./deployments/docker-compose.yml stop app

# Remove Go application container
remove: stop
	docker compose -f ./deployments/docker-compose.yml rm --force app

# Stop Go application
restart: stop remove start

# Remove Docker containers and generated images
clean: stop remove
	docker compose -f ./deployments/docker-compose.yml stop db
	docker compose -f ./deployments/docker-compose.yml rm --force db
	docker rmi -f ${APP_IMAGE} || true
	docker network rm ${NETWORK_NAME} || true
	docker images -f dangling=true -q | xargs docker rmi

# PostgreSQL Database ----------------------------------------
.PHONY: start-db stop-db remove-db clean-db	run-migrations rebuild-db

# Start the PostgreSQL database
start-db:
	@echo "Starting PostgreSQL database..."
	@echo "Pulling docker PostgreSQL image..."
	docker pull ${PGDB_IMAGE}
# TODO: implement persistance data(volumen)
	@echo "Running PostgreSQL database service..."
	docker-compose -f ./deployments/docker-compose.yml up -d ${PGDB_SERVICE_NAME}
	@echo "PostgreSQL database started."

# Stop the PostgreSQL database
stop-db:
	@echo "Stopping PostgreSQL database..."
	docker-compose -f ./deployments/docker-compose.yml stop ${PGDB_SERVICE_NAME}
	@echo "PostgreSQL database stopped."

# Remove the PostgreSQL database container
remove-db: stop-db
	@echo "Removing PostgreSQL database container..."
	docker-compose -f ./deployments/docker-compose.yml rm --force ${PGDB_SERVICE_NAME}
	@echo "PostgreSQL database container ${PGDB_SERVICE_NAME} removed."

# Clean the database 
clean-db: remove-db
	@echo "Cleaning PostgreSQL database image..."
	docker rmi -f ${PGDB_IMAGE} || true

# Run database migrations
run-migrations:
	@echo "Running database migrations..."
    # Add your migration commands here, e.g., using a migration tool like Flyway or a custom script
	@echo "Database migrations completed."

# Helper target for rebuilding the database (stop, clean, start)
rebuild-db: clean-db start-db
	@echo "Rebuilt PostgreSQL database."

# Show database logs
logs-db:
	@echo "Showing PostgreSQL database logs..."
	docker logs -f ${PGDB_CONTAINER_NAME}

# PgAdmin: PostgreSQL GUI ------------------------------------
.PHONY: pgadmin pgadmin-stop pgadmin-rm pgadmin-clean

pgadmin:
	@echo "Starting PgAdmin..."
	@echo "Pulling docker Pgadmin image..."
	docker pull ${PGADMIN_IMAGE}
# TODO: implement persistance data(volumen)
	@echo "Running PgAdmin service..."
	docker-compose -f ./deployments/docker-compose.yml up -d ${PGADMIN_SERVICE_NAME}
	@echo "PgAdmin started."

pgadmin-stop:
	@echo "Stopping PgAdmin..."
	docker compose -f ./deployments/docker-compose.yml stop ${PGADMIN_SERVICE_NAME}
	@echo "Stopped PgAdmin."

pgadmin-rm: pgadmin-stop
	@echo "Removing PgAdmin..."
	docker compose -f ./deployments/docker-compose.yml rm --force ${PGADMIN_SERVICE_NAME}
	@echo "Removed PgAdmin."

pgadmin-clean: pgadmin-rm
	docker rmi -f ${PGADMIN_IMAGE} || true

# Mocking ----------------------------------------------------
.PHONY: mocks
mocks:
	@echo "Running Mocks..."
	mockery --name=UserRepo --srcpkg=./internal/service --output=./internal/service/mocks
	mockery --name=UserService --srcpkg=./internal/controller --output=./internal/controller/mocks
	@echo "Mocks completed"

# Swagger Documentation --------------------------------------
.PHONY: swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/http-rest-api/main.go -o ./api

# Help -------------------------------------------------------
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  network ........ Create network for application"
	@echo "  app-img ........ Create and pull the required docker images"
	@echo "  build .......... Build the application"
	@echo "  app ............ Build and Run the Go application container"
	@echo "  start .......... Run Go application"
	@echo "  stop ........... Stop Go application"
	@echo "  remove ......... Stop and remove the Go application"
	@echo "  restart ........ Restart the Go application (stop,remove,start)"
	@echo "  clean .......... Stop, remove docker containers and generated images"
	@echo "  start-db ....... Start the PostgreSQL database"
	@echo "  stop-db ........ Stop the PostgreSQL database"
	@echo "  remove-db ...... Stop and remove the PostgreSQL database container"
	@echo "  clean-db ....... Stop and remove the PostgreSQL database container and docker image"
	@echo "  run-migrations . Run database migrations"
	@echo "  rebuild-db ..... Helper for rebuilding the database (stop, clean, start)"
	@echo "  logs-db ........ Show database logs"
	@echo "  pgadmin ........ Start the PgAdmin container"
	@echo "  pgadmin-stop ... Stop the PgAdmin container"
	@echo "  pgadmin-rm ..... Stop and remove the PgAdmin container"
	@echo "  pgadmin-clean .. Stop and remove the PgAdmin container and docker image"
	@echo "  mocks .......... Generate mock objects"
	@echo "  swagger ........ Generate swagger documentation"
