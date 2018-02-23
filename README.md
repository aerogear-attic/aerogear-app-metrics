# AeroGear Metrics API

This is the server component of the AeroGear Metrics Service. It is a RESTful API that allows mobile clients to send metrics data which will get stored in a PostgreSQL database. The service is written in [Golang](https://golang.org/).

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

See the [Contributing Guide](./CONTRIBUTING.md) for more information.


## Local development setup and building

This section documents how to setup a local development environment, if you only wish to run the project with minimal setup, check the [Containerized building and running](#containerized-building-and-running) section.

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

## Running the project

### Containerized building and running

You can use [Docker Compose](https://docs.docker.com/compose/) to run the project without requiring any extra setup other than a working docker installation.

```
docker-compose up
```

This will run both the `db` and `api` services. If you wish to run only the database use `docker-compose up db`.

The setup targets local development and testing, as such it utilizes the host's TCP ports 3000 for the API service and postgres' default port 5432.

This means that these ports could be in use if you have another postgres instance running or other test web servers.

### Running from source

Utilize the `go run` command to transparently compile and run the project:

```
go run cmd/metrics-api/metrics-api.go
```

You can verify the server is running by GET-ing its health check endpoint:
```
curl http://localhost:3000/healthz
```

The default configuration will allow the application to connect to the PostgreSQL container.

### Running a locally-built binary in a container

In order to run a locally-built binary in a container utilize the `Dockerfile.debug` file:

```
docker build -t aerogear/aerogear-metrics-api -f Dockerfile.debug .
```

This will copy a `./metrics` binary from the root of the project (the default output location for `make build` command),
and build an image from it.

NOTE: If not on a Linux platform, make sure to build the binary with `make build_linux` to use [go cross-compilation](http://golangcookbook.com/chapters/running/cross-compiling/) to build a Linux-compatible binary.

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

### Building a container image for a released binary

In order to build an image for a binary released with the previous process, utilize the `Dockerfile.prod` file:

```
docker build -t aerogear/aerogear-metrics-api -f Dockerfile.prod -build-arg version=0.0.x .
```

This will download the release's binary targetting the Linux OS to build the container image.

To publish it to your user's docker registry account:

```
# tag the image to your own namespace
docker tag aerogear/aerogear-metrics-api {your_username}/aerogear-metrics-api
docker login -u {your_username}
# your password is prompted
docker push {username}/aerogear-metrics-api
```
