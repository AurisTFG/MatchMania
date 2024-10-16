MAKEFLAGS += --no-print-directory

OUTPUT_DIR = bin
OUTPUT_FILE = MatchManiaAPI.exe
OUTPUT_PATH = $(OUTPUT_DIR)/$(OUTPUT_FILE)
SOURCE = main.go

.PHONY: all
all: run

.PHONY: build
build:
	@go build -o $(OUTPUT_PATH) $(SOURCE)
	@echo "Build completed"

.PHONY: run
run: build
	@echo "Server started"
	@$(OUTPUT_PATH)
	@echo "Server stopped"

.PHONY: rund # run with CompileDaemon
rund: build
	@echo "Server started"
	@CompileDaemon -command="$(OUTPUT_PATH)" -build="go build -o $(OUTPUT_PATH)"
	@echo "Server stopped"
	
.PHONY: migrate # migrate the database
migrate:
	@echo "Migrating started"
	@go run migrate/migrate.go
	@echo "Migrating ended"

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
	@echo "  make         - Builds and runs the application (default target)"
	@echo "  make build   - Builds the application"
	@echo "  make run     - Runs the application"
	@echo "  make rund    - Runs the application with CompileDaemon"
	@echo "  make migrate - Migrates the database"
	@echo "  make test    - Runs the tests"
	@echo "  make clean   - Cleans up the output directory"
	@echo "  make help    - Displays this help message"
