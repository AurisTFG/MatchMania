SHELL=/bin/bash
.DEFAULT_GOAL := help

.PHONY: install
install: ## Installs project dependencies
	@echo "[?] Installing project dependencies..."
	@pnpm install
	@echo "[✔] Dependencies installed"

.PHONY: build
build: ## Builds the project for production
	@echo "[?] Building the project for production..."
	@pnpm build
	@echo "[✔] Build completed"

.PHONY: run-dev
run-dev: install ## Starts the development server with live reloading
	@echo "[?] Starting the development server..."
	@pnpm dev
	@echo "[✔] Development server running"

.PHONY: run-prod
run-prod: install ## Previews the built application
	@echo "[?] Previewing the built application..."
	@pnpm preview
	@echo "[✔] Preview started"

.PHONY: lint
lint: ## Runs the linter to check for code style and errors
	@echo "[?] Running linter..."
	@pnpm lint:fix
	@echo "[✔] Linter completed"

.PHONY: test
test: ## Runs the project tests
	@echo "[?] Running tests..."
	@pnpm test:run
	@echo "[✔] Tests completed"

.PHONY: depcheck
depcheck: ## Checks for unused dependencies
	@echo "[?] Checking for unused dependencies..."
	@pnpm depcheck
	@echo "[✔] Dependency check completed"

.PHONY: prepare
prepare: install build lint depcheck test ## Prepares the code for pushing to the repository
	@echo "[✔] Code prepared for pushing to the repository"

.PHONY: clean
clean: ## Cleans up built files and dependencies
	@echo "[?] Cleaning up node_modules and dist folder..."
	@rm -rf node_modules dist
	@echo "[✔] Cleaned up node_modules and dist"

.PHONY: help
help: ## Displays this help message
	@echo "Makefile for MatchMania Frontend React Web Application"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-15s %s\n", $$1, $$2}'
	@echo ""
