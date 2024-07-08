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
