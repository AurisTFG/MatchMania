MAKEFLAGS += --no-print-directory

PACKAGE_MANAGER = pnpm

.PHONY: all
all: run-dev

.PHONY: install
install:
	@echo "[?] Installing project dependencies using $(PACKAGE_MANAGER)..."
	@$(PACKAGE_MANAGER) install
	@echo "[✔] Dependencies installed"

.PHONY: build
build: install
	@echo "[?] Building the project for production..."
	@$(PACKAGE_MANAGER) run build
	@echo "[✔] Build completed"

.PHONY: run-dev
run-dev: install
	@echo "[?] Starting the development server..."
	@$(PACKAGE_MANAGER) run dev
	@echo "[✔] Development server running"

.PHONY: run-prod
run-prod: build
	@echo "[?] Previewing the built application..."
	@$(PACKAGE_MANAGER) run preview
	@echo "[✔] Preview started"

.PHONY: test
test:
	@echo "[?] Running tests..."
	@$(PACKAGE_MANAGER) run test
	@echo "[✔] Tests completed"

.PHONY: lint
lint:
	@echo "[?] Running linter..."
	@$(PACKAGE_MANAGER) run lint
	@echo "[✔] Linter completed"

.PHONY: clean
clean:
	@echo "[?] Cleaning up node_modules and dist folder..."
	@rm -rf node_modules dist
	@echo "[✔] Cleaned up node_modules and dist"
	
.PHONY: prepare
prepare: lint build test
	@echo "[✔] Code prepared for pushing to the repository"

.PHONY: help
help:
	@echo "Makefile for MatchMania React Web Application"
	@echo ""
	@echo "Available commands:"
	@echo "  make          - Runs the development server (default target)"
	@echo "  make install  - Install project dependencies using $(PACKAGE_MANAGER)"
	@echo "  make build    - Build the project for production"
	@echo "  make run-dev  - Start the development server with live reloading"
	@echo "  make run-prod - Preview the built application"
	@echo "  make test     - Run the project tests"
	@echo "  make lint     - Run the linter to check for code style and errors"
	@echo "  make clean    - Clean up built files and dependencies"
	@echo "  make prepare  - Prepare the code for pushing to the repository"
	@echo "  make help     - Display this help message"
