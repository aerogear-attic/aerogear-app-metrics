# AeroGear Metrics API

**Please note this repo is a suggested implementation.** We have not agreed on implementing the service with Golang.

This is the server component of the AeroGear metrics service. It is a RESTful API that allows mobile clients to send metrics data which will get stored in a PostgreSQL database. The service is written in Golang.

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

## Clone and Install Dependencies

First clone this repository to `$GOPATH/src/github.com/aerogear/aerogear-metrics-api`

Then run the following command to install the dependencies

```
dep ensure
```

## How to Run

Use `docker-compose` to start the PostgreSQL container:

```
docker-compose up
```

Now you can build and run the application locally with the following command:

```
go run cmd/aerogear-metrics-api.go
```

The default configuration will allow the application to connect to the PostgreSQL container.

### How to Build

To build an executable of the application simply run:

```
go build cmd/aerogear-metrics-api
```

This will produce a binary called `aerogear-metrics-api`

### Docker Build

Simply run the following:

```
docker build -t aerogear/aerogear-metrics-api .
```