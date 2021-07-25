 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

VERSION       ?= $(shell git describe --tags --always --dirty)

all: fmt build test lint
build:
	$(GOBUILD) ./...
test:
	$(GOTEST) -short -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
lint:
	golangci-lint --timeout 1m run ./...
fmt:
	go fmt ./...
