SHELL=/bin/bash
.DEFAULT_GOAL := help

BUILD_FOLDER = ./bin
BUILD_PATH = $(BUILD_FOLDER)/MatchManiaAPI.exe
MAIN_FILE = ./cmd/server/main.go
COVERAGE_PATH = $(BUILD_FOLDER)/coverage.out

.PHONY: build
build: ## Builds the application
	@mkdir -p $(BUILD_FOLDER)
	@go build -o $(BUILD_PATH) $(MAIN_FILE)
	@echo "[✔] Build completed"

.PHONY: run-dev
run-dev: build ## Runs the application with CompileDaemon in Development environment
	@echo "[✔] Starting server with CompileDaemon in Development environment"
	@MATCHMANIA_ENV=dev CompileDaemon -command="$(BUILD_PATH)" -build="go build -o $(BUILD_PATH) $(MAIN_FILE)"
	@echo "[✔] Server stopped"

.PHONY: run-prod
run-prod: build ## Runs the application in Production environment
	@echo "[✔] Starting server in Production environment"
	@MATCHMANIA_ENV=prod $(BUILD_PATH)
	@echo "[✔] Server stopped"

.PHONY: lint
lint: ## Runs the linter to check for code style and errors
	@echo "[?] Running linter..."
	@golangci-lint run --enable-all
	@echo "[✔] Lint completed without errors"

.PHONY: tidy
tidy: ## Tidies up Go modules
	@echo "[?] Tidying up Go modules..."
	@go mod tidy
	@echo "[✔] Go modules tidied up"

.PHONY: test
test: ## Runs the tests
	@echo "[?] Running tests..."
	@go test ./... -v
	@echo "[✔] Tests completed"
	
.PHONY: coverage
coverage: ## Generates a coverage report
	@echo "[?] Generating coverage report..."
	@go test ./... -v -coverprofile=$(COVERAGE_PATH)
	@go tool cover -html=$(COVERAGE_PATH)
	@echo "[✔] Coverage report generated"

.PHONY: docs
docs: ## Generates API documentation with Swagger
	@echo "[?] Generating API documentation..."
	@swag init -g $(MAIN_FILE)
	@echo "[✔] API documentation generated"

.PHONY: clean
clean: ## Cleans up the output directory
	@echo "[?] Cleaning up output directory..."
	@rm -rf $(BUILD_FOLDER)
	@echo "[✔] Cleaned up output directory"

.PHONY: prepare
prepare: tidy test docs lint  ## Prepares the code for pushing to the repository
	@echo "[✔] Code prepared for pushing to the repository"

.PHONY: help
help: ## Displays this help message
	@echo "Makefile for MatchMania Backend Golang API Server"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-15s %s\n", $$1, $$2}'
	@echo ""
	