APP_NAME = aerogear-app-metrics
PKG     = github.com/aerogear/$(APP_NAME)
TOP_SRC_DIRS   = pkg
PACKAGES     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                   -exec dirname {} \\; | sort | uniq")
BIN_DIR := $(GOPATH)/bin
SHELL = /bin/bash
BINARY ?= metrics

# This follows the output format for goreleaser
BINARY_LINUX_64 = ./dist/linux_amd64/metrics

DOCKER_LATEST_TAG = aerogear/$(APP_NAME):latest
DOCKER_MASTER_TAG = aerogear/$(APP_NAME):master
RELEASE_TAG ?= $(CIRCLE_TAG)
DOCKER_RELEASE_TAG = aerogear/$(APP_NAME):$(RELEASE_TAG)

LDFLAGS=-ldflags "-w -s -X main.Version=${TAG}"

.PHONY: setup
setup:
	dep ensure

.PHONY: build
build: setup
	go build -o $(BINARY) ./cmd/metrics-api/metrics-api.go

.PHONY: build_linux
build_linux: setup
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX_64) ./cmd/metrics-api/metrics-api.go

.PHONY: docker_build
docker_build: build_linux
	docker build -t $(DOCKER_LATEST_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_build_release
docker_build_release:
	docker build -t $(DOCKER_LATEST_TAG) -t $(DOCKER_RELEASE_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_build_master
docker_build_master:
	docker build -t $(DOCKER_MASTER_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

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
		go test -tags=integration -coverprofile=coverage.out -covermode=count $(addprefix $(PKG)/,$(pkg)) || exit 1;\
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

.PHONY: docker_push_release
docker_push_release:
	@docker login --username $(DOCKERHUB_USERNAME) --password $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_LATEST_TAG)
	docker push $(DOCKER_RELEASE_TAG)
	
.PHONY: docker_push_master
docker_push_master:
	docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_MASTER_TAG)