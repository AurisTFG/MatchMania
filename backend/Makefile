MAKEFLAGS += --no-print-directory

OUTPUT_DIR = bin
OUTPUT_FILE_NAME = MatchManiaAPI.exe
SOURCE = src/main.go

.PHONY: all
all: run

.PHONY: build
build:
	@go build -o $(OUTPUT_DIR)/$(OUTPUT_FILE_NAME) $(SOURCE)
	@echo "Build completed"

.PHONY: run
run: build
	@echo "Server started"
	@./$(OUTPUT_DIR)/$(OUTPUT_FILE_NAME)
	@echo "Server stopped"

.PHONY: test
test:
	@go test -v ./...
	@echo "Test completed"

.PHONY: clean
clean:
	@rm -rf $(OUTPUT_DIR)
	@echo "Cleaned up $(OUTPUT_DIR)"

.PHONY: help
help:
	@echo "Makefile for MatchMania API Server"
	@echo "Available commands:"
	@echo "  make        - Builds and runs the application (default target)"
	@echo "  make build  - Builds the application"
	@echo "  make run    - Runs the application"
	@echo "  make test   - Runs the tests"
	@echo "  make clean  - Cleans up the output directory"
	@echo "  make help   - Displays this help message"
