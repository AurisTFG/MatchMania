.PHONY: all build run rund test lint clean gendocs help

MAKEFLAGS += --no-print-directory
OUTPUT_DIR = bin
OUTPUT_FILE = MatchManiaAPI.exe
OUTPUT_PATH = $(OUTPUT_DIR)/$(OUTPUT_FILE)
SOURCE = cmd/server/main.go

all: run

build:
	@go build -o $(OUTPUT_PATH) $(SOURCE)
	@echo "Build completed"

run: build
	@echo "Server started"
	@$(OUTPUT_PATH)
	@echo "Server stopped"

rund: build
	@echo "Server started with CompileDaemon"
	@MATCHMANIA_ENV=dev CompileDaemon -command="$(OUTPUT_PATH)" -build="go build -o $(OUTPUT_PATH) $(SOURCE)"
	@echo "Server stopped"

test: build
	@go test ./... -v -coverprofile $(OUTPUT_DIR)/coverage.out
	@go tool cover -html $(OUTPUT_DIR)/coverage.out
	@echo "Test completed"

lint:
	@go vet ./...
	@golangci-lint run
	@echo "Lint completed without errors"

gendocs:
	@swag init -g cmd/server/main.go
	@echo "Generated API documentation"

clean:
	@if [ -d $(OUTPUT_DIR) ]; then rm -rf $(OUTPUT_DIR); fi
	@echo "Cleaned up output directory \"$(OUTPUT_DIR)\""

help:
	@echo "Makefile for MatchMania API Server"
	@echo "Available commands:"
	@echo "  make         - Builds and runs the application (default target)"
	@echo "  make build   - Builds the application"
	@echo "  make run     - Runs the application"
	@echo "  make rund    - Runs the application with CompileDaemon"
	@echo "  make test    - Runs the tests"
	@echo "  make lint    - Runs the linter"
	@echo "  make gendocs - Generates the API documentation with Swagger"
	@echo "  make clean   - Cleans up the output directory"
	@echo "  make help    - Displays this help message"
