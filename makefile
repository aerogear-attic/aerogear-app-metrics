PKG     = github.com/aerogear/aerogear-metrics-api
TOP_SRC_DIRS   = pkg
TEST_DIRS     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                   -exec dirname {} \\; | sort | uniq")
BIN_DIR := $(GOPATH)/bin
SHELL = /bin/bash
BINARY = metrics
LINUX_BINARY = metrics_linux_amd_64
IMAGE_NAME = aerogear/aerogear-metrics-api

LDFLAGS=-ldflags "-w -s -X main.Version=${TAG}"

.PHONY: setup
setup:
	dep ensure

.PHONY: build
build: setup build_binary

.PHONY: build_binary_linux
build_binary_linux:
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY) ./cmd/metrics-api/metrics-api.go

.PHONY: build_binary
build_binary:
	go build -o $(BINARY) ./cmd/metrics-api/metrics-api.go

.PHONY: docker_build
docker_build: setup build_binary_linux
	docker build -t $(IMAGE_NAME) --build-arg BINARY=$(LINUX_BINARY)  -f deployments/docker/Dockerfile .

.PHONY: test-unit
test-unit:
	@echo Running tests:
	go test -v -race -cover $(UNIT_TEST_FLAGS) \
	  $(addprefix $(PKG)/,$(TEST_DIRS))

.PHONY: errcheck
errcheck:
	@echo errcheck
	@errcheck -ignoretests $$(go list ./...)

.PHONY: vet
vet:
	@echo go vet
	@go vet $$(go list ./...)

.PHONY: fmt
fmt:
	@echo go fmt
	diff -u <(echo -n) <(gofmt -d `find . -type f -name '*.go' -not -path "./vendor/*"`)

.PHONY: clean
clean:
	-rm -f ${BINARY}

.PHONY: release
release: setup
	goreleaser --rm-dist

.PHONY: build