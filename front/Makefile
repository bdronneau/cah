APP_NAME ?= front
VERSION ?= $(shell git log --pretty=format:'%h' -n 1)
AUTHOR ?= $(shell git log --pretty=format:'%an' -n 1)

MAKEFLAGS += --silent
GOBIN=bin
PID=/tmp/.$(APP_NAME).pid

.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sed -e 's|^| |'

.PHONY: $(APP_NAME)
## $(APP_NAME): Build app with dependencies download
$(APP_NAME): deps build

.PHONY: name
## name: Output name of app
name:
	@echo -n $(APP_NAME)

.PHONY: version
## version: Output sha1 of last commit
version:
	@echo -n $(VERSION)

.PHONY: author
## author: Output author's name of last commit
author:
	@python -c 'import sys; import urllib; print urllib.quote_plus(sys.argv[1])' "$(AUTHOR)"

.PHONY: deps
## deps: Install all needs for app
deps:
	npm install

.PHONY: build
## build: Build binary of app
build:
	NODE_ENV=production npm run build

.PHONY: start
## start: Start app
start:
	npm run serve