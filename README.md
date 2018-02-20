# AeroGear Metrics API

**Please note this repo is a suggested implementation.** We have not agreed on implementing the service with Golang.

This is the server component of the AeroGear metrics service. It is a RESTful API that allows mobile clients to send metrics data which will get stored in a PostgreSQL database. The service is written in Golang.

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

## Setup and Build

In Go, projects are typically kept in a [workspace](https://golang.org/doc/code.html#Workspaces) that follows a very specific architecture. Before cloning this repo, be sure you have a `GOPATH` env var that points to your workspace folder, say:

```sh
$ echo $GOPATH
/Users/JohnDoe/workspaces/go
```

Then clone this repository by running:

```
git clone git@github.com:aerogear/aerogear-metrics-api.git $GOPATH/src/github.com/aerogear/aerogear-metrics-api
```

And finally install dependencies:
```
make setup
```

It is also possible to build the binaries by simply running:
```
make build
```

## How to Run

In two different terminals:

1. Start the PostgreSQL container using `docker-compose up`:

```
docker-compose -f deployments/docker/docker-compose.yml up
```

2. Start the server app using `go run`:

```
go run cmd/metrics-api/metrics-api.go
```

You can verify it's running:
```
curl http://localhost:3000/healthz
```

The default configuration will allow the application to connect to the PostgreSQL container.

## Docker Build

Simply run the following:

```
cd deployments/docker
docker build -t aerogear/aerogear-metrics-api .
```

## Release

Build and publish to github releases using `goreleaser`, see `.goreleaser.yml` for configuration.

First make sure you have this tool installed: [Intalling Goreleaser](https://goreleaser.com/#introduction.installing_goreleaser).

Then tag your release, replacing `x` with the appropriate version:
```
git tag -a v0.0.x -m "Release 0.0.x"
```

And make the release:
```
make release
```
