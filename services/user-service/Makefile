.PHONY: all generate build run test clean help

ifeq ($(OS),Windows_NT)
    # Windows 系统
    SHELL = cmd
	SERVICE_PATH := $(shell echo %cd%)
	EXEC_DIR := $(SERVICE_PATH)\cmd
	# Binary output name
	BINARY_NAME=user_server.exe
else
    # Linux 系统
	SERVICE_PATH := $(shell pwd)
	EXEC_DIR := $(SERVICE_PATH)/cmd
	# Binary output name
	BINARY_NAME=user_server
endif

all: build

generate:
	@echo "Generating code..."
	go generate ./...

build:
	@echo "Building..."
	go build -ldflags="-w -s" -o cmd/$(BINARY_NAME) cmd/main.go

run:
	@echo "Running..."
	cd $(EXEC_DIR) && go run .

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	go clean
	rm -f cmd/$(BINARY_NAME)

# Help command to display available commands
help:
	@echo "Makefile commands:"
	@echo "all    - Build the application"
	@echo "build  - Build the binary"
	@echo "run    - Build and run the application"
	@echo "test   - Run tests"
	@echo "clean  - Remove binaries"
	@echo "help   - Display this help"

.DEFAULT_GOAL := build
