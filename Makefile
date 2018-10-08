# ########################################################## #
# Makefile for Golang Project
# Includes cross-compiling, installation, cleanup
# ########################################################## #

#The shell we use
SHELL := /bin/bash

# We like colors
# # From: https://coderwall.com/p/izxssa/colored-makefile-for-golang-projects
RED=`tput setaf 1`
GREEN=`tput setaf 2`
RESET=`tput sgr0`
YELLOW=`tput setaf 3`

# Go settings
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter
PATH_BUILD=test-pkgs

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


BINARY=ktl
VERSION := $(shell cat VERSION)
GIT_COMMIT := $(shell git rev-parse HEAD)
BUILD:= $(shell date -u +%FT%T)
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64
VERSION_PACKAGE = github.com/svx/ktl/pkg/version

# Setup linker flags option for build that interoperate with variable names in src code
#LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.GitCommit=$(GIT_COMMIT)"
LDFLAGS += -X main.Version=${VERSION}
LDFLAGS += -X main.GitCommit=${GIT_COMMIT}

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
.PHONY: help
help: ## This help message
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

.PHONY: default
default: build ## Builds default binary

all: clean build_all install

.PHONY: build
build: ## Build test binary
	@echo ""
	@echo "$(YELLOW)==> Creating test binaries for $(VERSION)$(RESET)"
	#go build ${LDFLAGS} -X github.com/svx/ktl/cmd/version -o $(PATH_BUILD)/$(BINARY)_$(VERSION)
	#go build -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}" -o $(PATH_BUILD)/$(BINARY)_$(VERSION)
	go build -ldflags "-X main.Version=${VERSION} -X github.com/svx/ktl/cmd.BuildTime=$(BUILD)" -o $(PATH_BUILD)/$(BINARY)_$(VERSION)
	#go build -ldflags '$(LDFLAGS)' -o $(PATH_BUILD)/$(BINARY)_$(VERSION)



build_gox:
	gox -ldflags  "-X main.Version=${VERSION} -X github.com/svx/ktl/cmd.BuildTime=$(BUILD)" -osarch="linux/amd64"

install:
	go install ${LDFLAGS}

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor
