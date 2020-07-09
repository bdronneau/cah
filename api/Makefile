SHELL = /bin/sh

APP_NAME ?= cah
VERSION ?= $(shell git log --pretty=format:'%h' -n 1)
AUTHOR ?= $(shell git log --pretty=format:'%an' -n 1)
PACKAGES ?= ./...

GOBIN=bin
BINARY_PATH=$(GOBIN)/$(APP_NAME)

## help: Display list of commands
.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sed -e 's|^| |'

## name: Output name
.PHONY: name
name:
	@echo -n $(APP_NAME)

## dist: Output binary path
.PHONY: dist
dist:
	@echo -n $(BINARY_PATH)

## version: Output sha1 of last commit
.PHONY: version
version:
	@echo -n $(VERSION)

## author: Output author's name of last commit
.PHONY: author
author:
	@python -c 'import sys; import urllib; sys.stdout.write(urllib.quote_plus(sys.argv[1]))' "$(AUTHOR)"

## $(APP_NAME): Build app with dependencies download
.PHONY: $(APP_NAME)
$(APP_NAME): deps go

## go: Build app
.PHONY: go
go: format lint tst build

## deps: Download dependencies
.PHONY: deps
deps:
	@echo -n $(GOPATH)
	go get github.com/kisielk/errcheck
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports

## format: Format code
.PHONY: format
format:
	goimports -w */*/*.go
	gofmt -s -w */*/*.go

## lint: Lint code
.PHONY: lint
lint:
	golint `go list $(PACKAGES) | grep -v vendor`
	# errcheck -verbose -ignoretests `go list $(PACKAGES) | grep -v vendor`
	# go vet $(PACKAGES)

## tst: Test with coverage
.PHONY: tst
tst:
	scripts/coverage || true

## build: Build binary
.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix nocgo -o $(BINARY_PATH) cmd/cah/main.go

## start: Start app
.PHONY: start
start:
	go run cmd/cah/main.go
