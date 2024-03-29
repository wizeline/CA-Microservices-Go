# Code Accelerator Golang Microservices (CAM-Go)
This project is a template generator based on Golang mainly focused on microservices. </br>
The purpose is to accelerate the development by providing a base solution avoiding starting from scratch.

## Requirements
- GO 1.22 or higher
- Docker
- Docker-Compose

## Architecture
High-Level Architecture Diagram. (to be defined)

## Features
- Hexagonal Architecture based
- Highly scalable
- Local-Dev: handy management tools through make statements
- GitHub Actions (CI/CD) base layout
- [Mockery](https://github.com/vektra/mockery) (mocks)
- [Viper](https://github.com/spf13/viper) support (env configuration)
- [Chi](https://github.com/go-chi/chi) (router)
- [ZeroLog](https://github.com/rs/zerolog) (logger)
- [PostgreSQL](https://www.postgresql.org/) database support
- [PgAdmin](https://www.pgadmin.org/) PostgreSQL database Web-GUI

## Generate Mocks with Mockery
`CAM-Go` applications use Mockery to generate mocks. This tool is able to handy generate and mantain your mock objects. It uses the stretchr/testify/mock package. Moreover, reduce the boilerplate code over mocking.

To add new mock objects please edit the [Makefile file](Makefile) look for the `mocks` target, and according to the Mockery's [documentation](https://github.com/vektra/mockery#readme) add the specific rule.

Before start, please make sure to [install](https://github.com/vektra/mockery#installation) `mockery` command line tool.

To automatically generate mocks, run `make mocks` at the top level of the project/repository.

## Local Dev
`CAM-Go` uses Docker and Docker-compose to generate a containerized environment with tools(make statements) to speed up development time.
Ensure you have configured the services environment variables in the `deployments/.env` file.

Run  `make help` to learn how to use all commands to handle your application environment.