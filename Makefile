# Makefile

# Define the binary output directory and name
BINARY_DIR := bin
BINARY_NAME := main
BINARY_PATH := $(BINARY_DIR)/$(BINARY_NAME)

# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build:
	@echo "Building the Go application..."
	@go build -o $(BINARY_PATH) .

# Run target
.PHONY: run
run: build
	@echo "Running the Go application..."
	@$(BINARY_PATH)

# Hot reload target using air
.PHONY: hot
hot:
	@echo "Starting the Go application with hot reloading using air..."
	@air

# Ensure Air is installed
.PHONY: ensure-air
ensure-air:
	@if ! [ -x "$$(command -v air)" ]; then \
		echo "Installing air..."; \
		go install github.com/cosmtrek/air@latest; \
	fi

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BINARY_DIR)

# Ensure all dependencies and tools are installed
.PHONY: setup
setup: ensure-air

# Default target including setup
.PHONY: default
default: setup all
