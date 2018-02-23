PKG     = github.com/aerogear/aerogear-metrics-api
TOP_SRC_DIRS   = pkg
PACKAGES     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                   -exec dirname {} \\; | sort | uniq")
BIN_DIR := $(GOPATH)/bin
SHELL = /bin/bash
BINARY ?= metrics

# This follows the output format for goreleaser
BINARY_LINUX_64 = ./dist/linux_amd64/metrics

TAG = aerogear/aerogear-metrics-api

LDFLAGS=-ldflags "-w -s -X main.Version=${TAG}"

.PHONY: setup
setup:
	dep ensure

.PHONY: build
build: | setup build_binary

.PHONY: build_linux
build_linux:
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX_64) ./cmd/metrics-api/metrics-api.go

.PHONY: build_binary
build_binary:
	go build -o $(BINARY) ./cmd/metrics-api/metrics-api.go

.PHONY: test-unit
test-unit:
	@echo Running tests:
	go test -v -race -cover $(UNIT_TEST_FLAGS) \
	  $(addprefix $(PKG)/,$(PACKAGES))

.PHONY: test-integration
test-integration:
	@echo Running tests:
	go test -v -race -cover $(UNIT_TEST_FLAGS) -tags=integration \
	  $(addprefix $(PKG)/,$(PACKAGES))

.PHONY: test-integration-cover
test-integration-cover:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -tags=integration -coverprofile=coverage.out -covermode=count $(addprefix $(PKG)/,$(pkg));\
		tail -n +2 coverage.out >> coverage-all.out;)

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

.PHONY: docker_build
docker_build: | setup build_linux
	docker build -t $(TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_untar_linux_release
docker_untar_linux_release:
	tar -xvf $(BINARY_LINUX_64).tar.gz

.PHONY: docker_release_build
docker_release_build: | setup release docker_untar_linux_release docker_build
	
.PHONY: docker_push
docker_push:
	docker push $(TAG)

.PHONY: build
