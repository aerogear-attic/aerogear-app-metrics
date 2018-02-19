# AeroGear Metrics API

**Please note this repo is a suggested implementation.** We have not agreed on implementing the service with Golang.

This is the server component of the AeroGear metrics service. It is a RESTful API that allows mobile clients to send metrics data which will get stored in a PostgreSQL database. The service is written in Golang.

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

## Clone and Install Dependencies

1. Clone this repository
```
go get github.com/aerogear/aerogear-metrics-api
```
1. Run the following command to build binary
```
make build
```

## How to Run

Use `docker-compose` to start the PostgreSQL container:

```
cd deployments/docker
docker-compose up
```

Now you can build and run the application locally with the following command:

```
go run cmd/metrics-api/metrics-api.go
```

The default configuration will allow the application to connect to the PostgreSQL container.

### Docker Build

Simply run the following:

```
cd deployments/docker
docker build -t aerogear/aerogear-metrics-api .
```

### Release

Builds and publish to github releases using `goreleaser`.
See `.goreleaser.yml` for configuration

Preparing release:
```
brew install goreleaser
## Replace x with version
git tag -a v0.0.X -m "Release 0.0.x"
make release
```