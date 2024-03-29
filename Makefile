
PROJECT := go-scaffold
VERSION := $(shell cat VERSION)
XC_OS 	:= linux darwin
XC_ARCH := 386 amd64 arm
LD_FLAGS := -X main.version=$(VERSION) -s -w
SOURCE_FILES ?=./...
TEST_PATTERN ?=.
TEST_OPTIONS ?=-v -failfast -race
TEST_TIMEOUT ?=2m
LINT_VERSION := 1.15.0

export XC_OS
export XC_ARCH
export VERSION
export PROJECT
export GO111MODULE := on
export LD_FLAGS
export SOURCE_FILES
export TEST_PATTERN
export TEST_OPTIONS
export TEST_TIMEOUT
export LINT_VERSION

.PHONY: all
all: help

.PHONY: help
help:
	@echo "make build - build $(PROJECT) for follwing OS-ARCH constilations: $(XC_OS) / $(XC_ARCH) "
	@echo "make build-dev - build $(PROJECT) for OS-ARCH set by GOOS and GOARCH env variables"
	@echo "make build-docker - build $(PROJECT) for linux-amd64 docker image"
	@echo "make fmt - use gofmt & goimports"
	@echo "make lint - run golangci-lint"
	@echo "make test - run go test including race detection"
	@echo "make coverage - same as test and uses go-junit-report to create report.xml"
	@echo "make dist - build and create packages with hashsums"
	@echo "make docker - creates a docker image"
	@echo "make docker-release/docker-release-latest - creates the docker image and pushes it to the registry (latest pushes also latest tag)"
	@echo "make setup - adds git pre-commit hooks"

.PHONY: build
build:
	@scripts/build.sh

.PHONY: build-dev
build-dev:
	@scripts/build.sh dev

.PHONY: build-docker
build-docker:
	@scripts/build.sh docker

.PHONY: dist
dist: 
	@scripts/dist.sh

.PHONY: fmt
fmt:
	@scripts/fmt.sh

.PHONY: lint
lint:
	@scripts/lint.sh

.PHONY: test
test:
	@scripts/test.sh

.PHONY: coverage
coverage:
	@scripts/test.sh coverage
	
.PHONY: docker
docker: build-docker
	@scripts/docker.sh

.PHONY: docker-release
docker-release:
	@scripts/docker.sh release

.PHONY: docker-release-latest
docker-release-latest:
	@scripts/docker.sh release latest

.PHONY: setup
setup: import
	@scripts/add-pre-commit.sh

.PHONY: import
import:
	go mod tidy

.PHONY: install
install:
	go install
	mv ${GOPATH}/bin/wikifier ${GOPATH}/bin/wikifier

go-git:
	git config --local url."git@github.com:".insteadOf "https://github.com/"

go-bitbucket:
	git config --local url."git@bitbucket.org:".insteadOf "https://bitbucket.org/" 

generate: 
	go run main.go generate --path=wiki --file=page1.md

apply: 
	go run main.go apply