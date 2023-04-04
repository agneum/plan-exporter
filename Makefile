.DEFAULT_GOAL = all

BINARY = plan-exporter

# Detect the OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    GOOS = linux
endif
ifeq ($(UNAME_S),Darwin)
    GOOS = darwin
endif

# Detect the architecture
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
    GOARCH = amd64
endif
ifeq ($(UNAME_M),arm64)
    GOARCH = arm64
endif
ifeq ($(UNAME_M),aarch64)
    GOARCH = arm64
endif

VERSION?=0.0.1
BUILD_TIME?=$(shell date -u '+%Y%m%d-%H%M')
COMMIT?=$(shell git rev-parse HEAD)
BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)

# Set up GOPATH if not defined
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

# Symlink into GOPATH
BUILD_DIR=${GOPATH}/${BINARY}

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w \
	-X main.version=${VERSION} \
	-X main.commit=${COMMIT} \
	-X main.branch=${BRANCH}\
	-X main.buildTime=${BUILD_TIME}"

# Go tooling command aliases
GOBUILD = GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS}
GOTEST = GO111MODULE=on go test
GORUN = GO111MODULE=on go run ${LDFLAGS}

# Build the project
all: clean build

# Install the linter to $GOPATH/bin which is expected to be in $PATH
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin v1.52.2

run-lint:
	golangci-lint --config .golangci.yml run

lint: install-lint run-lint

build:
	${GOBUILD} -o bin/${BINARY}-${GOOS}-${GOARCH} ./main.go

clean:
	-rm -f bin/*

run:
	${GORUN} ./main.go

test:
	${GOTEST} -v ./...

test-race:
	${GOTEST} -v -race ./...

install: build
	cp bin/${BINARY}-${GOOS}-${GOARCH} /usr/local/bin/${BINARY}

uninstall:
	rm -f /usr/local/bin/${BINARY}

.PHONY: all clean build lint run-lint install-lint test test-race install uninstall
