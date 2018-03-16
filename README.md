# Aerogear App Metrics

[![Go Report Card](https://goreportcard.com/badge/github.com/aerogear/aerogear-app-metrics)](https://goreportcard.com/report/github.com/aerogear/aerogear-app-metrics)
[![Coverage Status](https://coveralls.io/repos/github/aerogear/aerogear-app-metrics/badge.svg?branch=master)](https://coveralls.io/github/aerogear/aerogear-app-metrics?branch=master)
[![CircleCI](https://circleci.com/gh/aerogear/aerogear-app-metrics.svg?style=svg)](https://circleci.com/gh/aerogear/aerogear-app-metrics)

This is the server component of the AeroGear Metrics Service. It is a RESTful API that allows mobile clients to send metrics data which will get stored in a PostgreSQL database. The service is written in [Golang](https://golang.org/).

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Ensure the $GOPATH environment variable is set](https://github.com/golang/go/wiki/SettingGOPATH)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

See the [Contributing Guide](./CONTRIBUTING.md) for more information.


## Getting Started

This section documents the ideal setup for local development. If you'd like to simply run the entire application in `docker-compose`, check out the relevant section.

Golang projects are kept in a [workspace](https://golang.org/doc/code.html#Workspaces) that follows a very specific architecture. Before cloning this repo, be sure you have a `$GOPATH` environment variable set up.

### Clone the Repsitory

```
git clone git@github.com:aerogear/aerogear-app-metrics.git $GOPATH/src/github.com/aerogear/aerogear-app-metrics
```

### Install Dependencies
```
make setup
```

Note this is using the `dep` package manager under the hood. You will see the dependencies installed in the `vendor` folder.

### Start the Database

```
docker-compose up db
```

### Start the Server

```
go run cmd/metrics-api/metrics-api.go

2018/02/27 10:51:54 Starting application... going to listen on :3000
```

The application's default configuration will allow it to connect to the database created by `docker-compose`.

You can test out the healthcheck endpoint:

```
curl -i http://localhost:3000/healthz

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Tue, 27 Feb 2018 10:53:16 GMT
Content-Length: 56

{"time":"2018-02-27T10:53:16.313301415Z","status":"ok"}
```

### Run Entire Application with Docker Compose

This section shows how to start the entire application with `docker-compose`. This is useful for doing some quick tests (using the SDKs) for example.

First, compile a Linux compatible binary:

```
make build_linux
```

This binary will be used to build the Docker image. Now start the entire application.

```
docker-compose up
```

### Example Client Request

This section shows an example `curl` request which can be used to send some data to the `/metrics` endpoint.

```
curl -i -X POST \
  http://localhost:3000/metrics \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 87bf2b99-7cdc-8df9-9b2d-6cdcd2932159' \
  -d '{
  "clientId": "453de7432",
  "data": {
    "app": {
      "appId": "com.example.someApp",
      "sdkVersion": "2.4.6",
      "appVersion": "256"
    },
    "device": {
      "platform": "android",
      "platformVersion": "27"
    }
  }
}'
```

You will see the data returned back in the response. 

If you have the `psql` command line tools you can connect to the Database and verify the data was inserted.

```
psql -U postgresql -d aerogear_mobile_metrics --host localhost
Password for user postgresql: # postgres

aerogear_mobile_metrics=> select * from mobileappmetrics;
```

## Builds and Testing

The `makefile` provided provides commands for building and testing the code. For now, only the most important commands are documented.

* `make setup` - Downloads the application dependencies.

* `make build` - Compiles a binary for your current system. The binary is output at `./aerogear-app-metrics`.

* `make build_linux` - Compiles a Linux compatible binary. The binary is output at `./dist/linux_amd64/aerogear-app-metrics`. This is mainly used for Docker builds. `make build` can still be used if you are a Linux user.

* `make docker_build` - Builds a Binary from source and uses that binary to create a Docker image.

* `make test` - Runs the unit tests (alias for `make test-unit`).

* `make test-integration` - Runs unit tests and integration tests. This requires a running database server.

* `make test-integration-cover` - Same as `make test-integration` but also generates a code coverage report. Used in the CI service.

## Environment Variables

The aerogear-app-metrics server is configured using environment variables:

### Server Configuration

| Variable          | Default | Description                                                                |
|-------------------|---------|----------------------------------------------------------------------------|
| PORT              | 3000    | The port the server will listen on                                         |
| LOG_LEVEL         | info    | Can be one of [debug, info, warning, error, fatal, panic]                  |
| LOG_FORMAT        | text    | Can be one of [text, json]                                                 |
| DBMAX_CONNECTIONS | 100     | The maximum number of concurrent database connections the server will open |

### Database Connection Parameters

The database connection is configured using the table of environment variables below. These environment variables correspond to the PostgreSQL [libpq environment variables](https://www.postgresql.org/docs/current/static/libpq-envars.html). The table below shows all of the environment variables supported by the `pq` driver used in this server.

| Variable          | Default                 | Description                                                                                                                                   |
|-------------------|-------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| PGDATABASE        | aerogear_mobile_metrics | The database to connect to                                                                                                                    |
| PGUSER            | postgresql              | The database user                                                                                                                             |
| PGPASSWORD        | postgres                | The database password                                                                                                                         |
| PGHOST            | localhost               | The database hostname to connect to                                                                                                           |
| PGPORT            | 5432                    | The database port to connect to                                                                                                               |
| PGSSLMODE         | disable                 | The SSL mode                                                                                                                                  |
| PGCONNECT_TIMEOUT | 5                       | The default connection timeout (seconds)                                                                                                      |
| PGAPPNAME         |                         | The [application_name](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNECT-APPLICATION-NAME) connection parameter |
| PGSSLCERT         |                         | The [sslcert](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNECT-SSLCERT) connection parameter.                  |
| PGSSLKEY          |                         | The [sslkey](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNECT-SSLKEY) connection parameter.                    |
| PGSSLROOTCERT     |                         | The [sslrootcert](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNECT-SSLROOTCERT) connection parameter           |


## How to Release a New Version

The release process for Aerogear maintainers is very simple. From the Github UI, simply create a release. **The release tag and title must be in the format `x.y.z`**. Formats such as `1.0` or `v1.0.0` are not valid.

This will kick off an automated process in the CI service, where the following steps occur:

* The release tag is checked out.
* The code is built and the full test suite is run.
* The [goreleaser](https://goreleaser.com/) tool generates a Changelog and binaries for Windows, MacOS and Linux. These are added to the release created in Github.
* A new docker image is built and given the tags `latest` and `<git tag>` (where `<git tag>` is the `x.y.z` tag that was used).
* The docker image is pushed to the Aerogear organization in Github.

The automated release process takes 2-3 minutes to complete on average.

copyright AeroGear 2018.
