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

In order to develop the Metrics API, its source code should be contained inside your `$GOPATH` and in the proper directory under `$GOPATH/src/`:

```bash
git clone https://github.com/aerogear/aerogear-app-metrics.git $GOPATH/src/github.com/aerogear/aerogear-app-metrics
cd aerogear-app-metrics/
```

See the [Go wiki](https://github.com/golang/go/wiki/GOPATH) for more information.

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
3. Follow the Commit Message Guidelines

## Commit Message Guidelines

This project has rules for commit messages (based on [Conventional Commits](https://conventionalcommits.org/)). There are several reasons for this:

* Commit messages are more readable, especially when looking through the **project history**.
* Commit messages describe whether a major, minor or patch change has been introduced (see [semver.org](https://semver.org/))
* Commit messages can be used to generate a changelog.

### Commit Message Format
Each commit message consists of a **header**, an optional **body** and an an optional **footer**.  The header has a special
format that includes a **type**, an optional **scope** and a **subject**:

```
<type>(optional scope): <description>

[optional body]

[optional footer]
```

99% of the time, you'll simply add a prefix so your message looks like one of these three:

```
fix: <your message>
```

```
feat: <your message>
```

```
breaking: <your message>

<description of breaking change>
```

Where `fix` represents a semver patch change, `feat` represents a semver minor change and `breaking` represents a semver major change.

Here are some examples:

```
fix: clean up test commands in makefile
```

```
feat: initial parsing of security metrics (#33)
```

```
breaking: rename Type field in SecurityMetric to Id (#37)

The security metric must contain an `id` field instead `type`
```

Try to keep each line of the commit message shorter than 100 characters. This allows the message to be easier
to read on GitHub as well as in various git tools.

### Type
In most cases you will use one of these three:

* **fix**: A bug fix
* **feat**: A new feature
* **breaking**: A breaking change, should also include a description of the change in the footer.

**It is not required**, but you can you can also choose from one of the following if you think one is more appropriate:

* **build**: Changes that affect the build system or external dependencies
* **ci**: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, Jenkins)
* **docs**: Documentation only changes
* **perf**: A code change that improves performance
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **test**: Adding missing tests or correcting existing tests

### Scope
The scope is optional, but generally it should be the name of the internal package or component affected. Some examples:

```
fix(dao) use new config and set max db connections
```

```
fix(config): major improvements to config parser
```

Where `dao` and `config` are internal packages.

### Footer
The footer should contain any information about **Breaking Changes** and is also the place to
reference issues (JIRA or Github) that this commit closes.

### Revert
If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit. In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.

