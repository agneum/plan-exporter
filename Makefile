.DEFAULT_GOAL = all

BINARY = plan-exporter
GOARCH = amd64
GOOS = linux

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

# Build the project
all: clean build

# Install the linter to $GOPATH/bin which is expected to be in $PATH
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin v1.23.8

run-lint:
	golangci-lint --config .golangci.yml run

lint: install-lint run-lint

build:
	${GOBUILD} -o bin/${BINARY} ./cmd/joe/main.go

clean:
	-rm -f bin/*

run:
	go run ${LDFLAGS} ./cmd/joe/main.go

.PHONY: all clean build lint run-lint install-lint

