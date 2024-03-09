# Code Accelerator Golang Microservices
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
- [Echo](https://echo.labstack.com/) Framework (http-server)
- [ZeroLog](https://github.com/rs/zerolog) (logger)
- [Mockery](https://github.com/vektra/mockery) (mocks)
- GitHub Actions (CI/CD) 
- Highly scalable
- Handy management tools through make statements

## Generate Mocks with Mockery
`CA-Microservices-Go` applications use Mockery to generate mocks. This tool is able to handy generate and mantain your mock objects. It uses the stretchr/testify/mock package. Moreover, reduce the boilerplate code over mocking.

To add new mock objects please edit the [Makefile file](Makefile) look for the `mocks` target, and according to the Mockery's [documentation](https://github.com/vektra/mockery#readme) add the specific rule.

Before start, please make sure to [install](https://github.com/vektra/mockery#installation) `mockery` command line tool.

To automatically generate mocks, run `make mocks` at the top level of the project/repository.
