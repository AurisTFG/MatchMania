MAKEFLAGS += --no-print-directory

BUILD_FOLDER = ./bin
BUILD_PATH = $(BUILD_FOLDER)/MatchManiaAPI.exe
MAIN_FILE = ./cmd/server/main.go
COVERAGE_PATH = $(BUILD_FOLDER)/coverage.out

.PHONY: all
all: run-dev

.PHONY: build
build:
	@mkdir -p $(BUILD_FOLDER)
	@go build -o $(BUILD_PATH) $(MAIN_FILE)
	@echo "[✔] Build completed"

.PHONY: run-dev
run-dev: build
	@echo "[✔] Starting server with CompileDaemon in Development environment"
	@MATCHMANIA_ENV=dev CompileDaemon -command="$(BUILD_PATH)" -build="go build -o $(BUILD_PATH) $(MAIN_FILE)"
	@echo "[✔] Server stopped"

.PHONY: run-prod
run-prod: build
	@echo "[✔] Starting server in Production environment"
	@MATCHMANIA_ENV=prod $(BUILD_PATH)
	@echo "[✔] Server stopped"

.PHONY: test
test: build
	@echo "[?] Running tests..."
	@go test ./...
	@echo "[✔] Tests completed"

.PHONY: coverage
coverage: build
	@echo "[?] Generating coverage report..."
	@go test ./... -v -coverprofile=$(COVERAGE_PATH)
	@go tool cover -html=$(COVERAGE_PATH)
	@echo "[✔] Coverage report generated"

.PHONY: docs
docs:
	@echo "[?] Generating API documentation..."
	@swag init -g $(MAIN_FILE) > /dev/null 2>&1
	@echo "[✔] API documentation generated"

.PHONY: vendor
vendor:
	@echo "[?] Creating vendor directory..."
	@go mod vendor
	@echo "[✔] Vendor directory created"

.PHONY: vet
vet:
	@echo "[?] Running Go vet..."
	@go vet ./...
	@echo "[✔] Go vet completed"

.PHONY: lint
lint:
	@echo "[?] Running linter..."
	@golangci-lint run --enable-all
	@echo "[✔] Lint completed without errors"

.PHONY: tidy
tidy:
	@echo "[?] Tidying up Go modules..."
	@go mod tidy
	@echo "[✔] Go modules tidied up"

.PHONY: format
format:
	@echo "[?] Formatting code..."
	@go fmt ./...
	@echo "[✔] Code formatted"

.PHONY: clean
clean:
	@echo "[?] Cleaning up output directory..."
	@rm -rf $(BUILD_FOLDER)
	@echo "[✔] Cleaned up output directory"

.PHONY: prepare
prepare: format vet tidy clean build test docs
	@echo "[✔] Code prepared for pushing to the repository"

.PHONY: help
help:
	@echo "Makefile for MatchMania API Server"
	@echo ""
	@echo "Available commands:"
	@echo "  make          - Builds and runs the application (default target)"
	@echo "  make build    - Builds the application"
	@echo "  make run-dev  - Runs the application with CompileDaemon in Development environment"
	@echo "  make run-prod - Runs the application in Production environment"
	@echo "  make test     - Runs the tests"
	@echo "  make coverage - Generates a coverage report"
	@echo "  make docs     - Generates the API documentation with Swagger"
	@echo "  make vendor   - Creates a vendor directory with dependencies"
	@echo "  make vet      - Runs Go vet, a static analysis tool"
	@echo "  make lint     - Runs the linter, checks for code style and errors"
	@echo "  make tidy     - Tidies up Go modules"
	@echo "  make format   - Formats the code"
	@echo "  make clean    - Cleans up the output directory"
	@echo "  make prepare  - Prepares the code for pushing to the repository"
	@echo "  make help     - Displays this help message"
