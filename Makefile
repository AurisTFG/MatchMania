MAKEFLAGS += --no-print-directory

BACKEND_FOLDER = ./backend/
FRONTEND_FOLDER = ./frontend/

.PHONY: all
all: run-dev

.PHONY: run-dev
run-dev:
	@$(MAKE) -C $(BACKEND_FOLDER) run-dev &
	@$(MAKE) -C $(FRONTEND_FOLDER) run-dev &

.PHONY: run-prod
run-prod:
	@$(MAKE) -C $(BACKEND_FOLDER) run-prod &
	@$(MAKE) -C $(FRONTEND_FOLDER) run-prod &

.PHONY: prepare
prepare: 
	@$(MAKE) -C $(BACKEND_FOLDER) prepare
	@$(MAKE) -C $(FRONTEND_FOLDER) prepare

.PHONY: help
help:
	@echo "Master Makefile for React and Go Projects"
	@echo ""
	@echo "Available commands:"
	@echo "  make         - Runs both the frontend and backend development servers"
	@echo "  make run-front - Starts the React development server"
	@echo "  make run-back  - Starts the Go development server"
	@echo "  make prepare   - Prepares both frontend and backend projects"
	@echo "  make help      - Displays this help message"
