# Golang Boilerplate API

![Logo](./.github/GopportunitiesHeader.svg) 

This is a job opportunities API built using the Golang programming language. The API uses the Gin framework as a router, SQLite as a database, and GoORM for database communication. It also includes Swagger for documentation and testing, and a well-organized package structure.


## Prerequisites

To run this project, you will need to have installed on your machine:

- Golang
- SQLite3
- Swarggo


## Features

- Introduction to Golang and building modern APIs
- Development environment setup for creating the API
- Using Go-Gin as a router for route management
- Implementing SQLite as the database for the API
- Using GoORM for communication with the database
- Integrating Swagger for API documentation and testing
- Creating a modern package structure for organizing the project
- Implementing a complete job opportunities API with endpoints for searching, creating, editing, and deleting opportunities
- Implementing automated tests to ensure API quality


## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/LuisMarchio03/golang-boilerplate-api`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `golang-boilerplate-api`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `golang-boilerplate-api` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

For more information on Docker and Docker Compose, refer to the official documentation:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)


The server will be running at http://localhost:8080.

## Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing
## Documentation

The API documentation is available at http://localhost:8080/swagger/index.html. Here, you will find all available endpoints, their parameters, and examples of requests and responses.

## License

This project is licensed under the MIT License - see the [LICENSE](https://choosealicense.com/licenses/mit/) file for details.



