# Contributing to the AeroGear Metrics API

The AeroGear Metrics API is part of the [AeroGear project](https://aerogear.org/), see the [Community Page](https://aerogear.org/community) for general guidelines for contributing to the project.

This document details specifics for contributions to the Metrics API.

## Issue tracker

The tracking of issues for the AeroGear Metrics API is done in the [AeroGear Android Project](https://issues.jboss.org/projects/AEROGEAR/issues) in the [JBoss Developer JIRA](https://issues.jboss.org).

See the [AeroGear JIRA Usage and Guidelines Guide](https://aerogear.org/docs/guides/JIRAUsage/) for information on how the issue tracker relates to contributions to this project.

## Asking for help

Whether you're contributing a new feature or bug fix, or simply submitting a
ticket, the Aerogear team is available for technical advice or feedback. 
You can reach us at [#aerogear](ircs://chat.freenode.net:6697/aerogear) on [Freenode IRC](https://freenode.net/) or the 
[aerogear google group](https://groups.google.com/forum/#!forum/aerogear)
-- both are actively monitored.

# Developing the AeroGear Metrics API

## Prerequisites

Ensure you have the following installed in your machine:

- [Golang](https://golang.org/dl/)
- [GNU Make](https://www.gnu.org/software/make/)
- [Git SCM](http://git-scm.com/)

## Cloning the repository

```bash
git clone https://github.com/aerogear/aerogear-metrics-api.git
cd aerogear-metrics-api/
```

## Installing dependencies and building the Metrics API

Run the following commands to install the dependencies for the Metrics API and build the resulting binary:

```bash
# Install the `dep` dependency manager
go get github.com/golang/dep/cmd/dep
make build
```

Refer to the [README](./README.md) for more details and alternative building utilizing containers.

## Before submitting a Pull Request

1. Make sure unit tests are passing locally with `make test-unit`
2. Format the source code with `make fmt`
