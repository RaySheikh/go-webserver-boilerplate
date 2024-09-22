# Define variables
APP_NAME := go-webserver-boilerplate
CONFIG_FILE := config/config.yaml
MAIN_FILE := main.go
OUTPUT_DIR := bin
IMAGE_NAME := go-webserver
DOCKERFILE := Dockerfile

# Default port (overridden by config)
PORT := 8080

# Targets
.PHONY: all run build test clean lint docker-build docker-run docker-clean

# Default to build
default: build

## run: Runs the Go web server
run: $(CONFIG_FILE)
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

## build: Builds the Go web server binary
build: $(CONFIG_FILE)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(APP_NAME) $(MAIN_FILE)

## test: Runs all tests
test:
	@echo "Running tests..."
	@go test ./...

## clean: Removes the build binary and cached Go files
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@go clean

## lint: Runs static analysis using golangci-lint (you need to install golangci-lint)
lint:
	@echo "Running linter..."
	@golangci-lint run

## docker-build: Builds the Docker image
docker-build:
	@echo "Building Docker image $(IMAGE_NAME)..."
	@docker build -t $(IMAGE_NAME) .

## docker-run: Runs the Docker container
docker-run: docker-build
	@echo "Running Docker container $(IMAGE_NAME)..."
	@docker run -p $(PORT):8080 $(IMAGE_NAME)

## docker-clean: Removes the Docker image
docker-clean:
	@echo "Cleaning up Docker images and containers..."
	@docker ps -q --filter "ancestor=$(IMAGE_NAME)" | xargs -r docker stop
	@docker ps -a -q --filter "ancestor=$(IMAGE_NAME)" | xargs -r docker rm
	@docker rmi -f $(IMAGE_NAME) || true


