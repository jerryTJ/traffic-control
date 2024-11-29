# Makefile for Go project

# 变量定义
APP_NAME = proxy-traffic
SRC_DIR = .
BIN_DIR = ./bin
GO_FILES = $(wildcard $(SRC_DIR)/*.go)
TEST_FILES = $(wildcard $(SRC_DIR)/*_test.go)
GOOS ?= $(shell go env GOOS)              
GOARCH ?= $(shell go env GOARCH)          
VERSION  := 1.0.0 
IMAGE_NAME := liuzhidocker/proxy-traffic
DOCKERFILE := Dockerfile

# default 
all: build

# clean
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)/$(APP_NAME)

# download dependency
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod tidy

# compile
.PHONY: build
build: $(BIN_DIR)/$(APP_NAME)

$(BIN_DIR)/$(APP_NAME): $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)/Application.go

# run 
.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	$(BIN_DIR)/$(APP_NAME)

# format code
.PHONY: fmt
fmt:
	@echo "Checking code formatting..."
	gofmt -l $(GO_FILES)

# run test
.PHONY: test
test:
	@echo "Running tests..."
	go test $(SRC_DIR)

# 
.PHONY: lint
lint:
	@echo "Running linter..."
	golangci-lint run $(SRC_DIR)

# build docker image
.PHONY: docker-build
docker-build: build
	@echo "Building Docker image $(IMAGE_NAME):$(VERSION)..."
	docker build -t $(IMAGE_NAME):$(VERSION) -f $(DOCKERFILE) .

#  push image to hub
.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image $(IMAGE_NAME):$(VERSION)..."
	docker push $(IMAGE_NAME):$(VERSION)

# run all target
.PHONY: all
all: clean deps fmt lint test build run 
