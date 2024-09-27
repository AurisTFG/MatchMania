MAKEFLAGS += --no-print-directory

PACKAGE_MANAGER = pnpm

INSTALL_CMD   = $(PACKAGE_MANAGER) install
BUILD_CMD     = $(PACKAGE_MANAGER) run build
RUN_CMD       = $(PACKAGE_MANAGER) run dev
PREVIEW_CMD   = $(PACKAGE_MANAGER) run preview
TEST_CMD      = $(PACKAGE_MANAGER) run test
CLEAN_CMD     = rm -rf node_modules dist

.PHONY: all
all: run

.PHONY: install
install:
	@echo "Installing project dependencies using $(PACKAGE_MANAGER)"
	@$(INSTALL_CMD)

.PHONY: build
build: install
	@echo "Building the project"
	@$(BUILD_CMD)

.PHONY: run
run: install
	@echo "Starting the development server"
	@$(RUN_CMD)

.PHONY: preview
preview: build
	@echo "Previewing the built application"
	@$(PREVIEW_CMD)

.PHONY: test
test:
	@$(TEST_CMD)
	@echo "Test completed"
	
.PHONY: clean
clean:
	@$(CLEAN_CMD)
	@echo "Cleaned up node_modules and dist"

.PHONY: help
help:
	@echo "Makefile for MatchMania React App"
	@echo "Available commands:"
	@echo "  make          - Runs the development server (default target)"
	@echo "  make install  - Install project dependencies using $(PACKAGE_MANAGER)"
	@echo "  make build    - Build the project for production"
	@echo "  make run      - Start the development server with live reloading"
	@echo "  make preview  - Preview the built application"
	@echo "  make test     - Run the project tests"
	@echo "  make clean    - Clean up built files and dependencies"
	@echo "  make help     - Display this help message"
