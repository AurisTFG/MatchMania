SHELL=/bin/bash
.DEFAULT_GOAL := help

BACKEND_FOLDER = ./backend/
FRONTEND_FOLDER = ./frontend/
MAKE_CMD = @$(MAKE) -C

.PHONY: run-db
run-db: ## Starts the postgreSQL database server
	pg_ctl start -D "C:\Program Files\PostgreSQL\17\data"

.PHONY: run-back
run-back: ## Starts the Go development server
	$(MAKE_CMD) $(BACKEND_FOLDER) run-dev

.PHONY: run-front
run-front: ## Starts the React development server
	$(MAKE_CMD) $(FRONTEND_FOLDER) run-dev

.PHONY: prepare
prepare: ## Prepares the code for pushing to the repository
	$(MAKE_CMD) $(BACKEND_FOLDER) prepare
	$(MAKE_CMD) $(FRONTEND_FOLDER) prepare

.PHONY: clean
clean: ## Cleans up temporary files in both frontend and backend
	$(MAKE_CMD) $(BACKEND_FOLDER) clean
	$(MAKE_CMD) $(FRONTEND_FOLDER) clean

.PHONY: cloc
cloc: ## Counts lines of code in both frontend and backend
	@cloc $(BACKEND_FOLDER) $(FRONTEND_FOLDER) --by-file-by-lang --exclude-lang=YAML,JSON,make,Dockerfile,SVG,HTML,CSS,JavaScript --exclude-dir=bin,dist,node_modules,docs

.PHONY: cloc-test
cloc-test: ## Counts lines of code in both frontend and backend, including test files
	@cloc $(BACKEND_FOLDER) $(FRONTEND_FOLDER) --by-file-by-lang --exclude-lang=YAML,JSON,make,Dockerfile,SVG,HTML,CSS,JavaScript --exclude-dir=bin,dist,node_modules,docs --not-match-f='.*test.*'

.PHONY: help
help: ## Displays this help message
	@echo "Master Makefile for React and Go Projects"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-15s %s\n", $$1, $$2}'
	@echo ""
