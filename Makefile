# AI30 AirTraffic Makefile
# Author: Victor Dessenne
# Date: Décembre 2025
# Updated:
# This Makefile is used to build and manage the AI30 AirTraffic project.

.PHONY: help dev build-frontend build-backend build-tools prod run-prod clean clean-data install generate-data check-prereqs

# Shell configuration for portability
SHELL := /bin/bash
.SHELLFLAGS := -ec

# Project root directory (absolute path)
ROOT_DIR := $(shell pwd)

# Color codes for terminal output
BLUE := \033[0;34m
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m # No Color

# Project paths
FRONTEND_DIR := frontend/simulation-front-end
BACKEND_DIR := backend/
BINARY_NAME := air-traffic-simulator
DATABASE_DIR := db
BINARY_DIR := bin

# Tool binaries
GET_AIRPLANES_BIN := $(BINARY_DIR)/get-airplanes-data
DOWNLOAD_ROUTES_BIN := $(BINARY_DIR)/download-routes
XML2JSON_BIN := $(BINARY_DIR)/xml2json

# Ports
BACKEND_PORT := 7500
FRONTEND_PORT := 3000

# Options
KEEP_SNAPSHOT ?= no  # Options: yes ou no (défaut: no - renouvelle le snapshot)

# ============================================================================
# HELP TARGET (runs first if you just type 'make')
# ============================================================================
help:
	@echo -e "$(BLUE)╔════════════════════════════════════════════════════════════╗$(NC)"
	@echo -e "$(BLUE)║         AI30 AirTraffic - Available Commands               ║$(NC)"
	@echo -e "$(BLUE)╚════════════════════════════════════════════════════════════╝$(NC)"
	@echo -e ""
	@echo -e "$(GREEN)Development (with hot reload):$(NC)"
	@echo -e "  $(YELLOW)make dev$(NC)              Start React (3000) and Go API (7500) in separate terminals"
	@echo -e "  $(YELLOW)make dev-frontend$(NC)     Start only React dev server on port 3000"
	@echo -e "  $(YELLOW)make dev-backend$(NC)      Start only Go API on port 7500 (proxies React on 3000)"
	@echo -e ""
	@echo -e "$(GREEN)Production (single executable):$(NC)"
	@echo -e "  $(YELLOW)make prod$(NC)             Build single executable (builds frontend + backend)"
	@echo -e "  $(YELLOW)make run-prod$(NC)         Run the production executable on port 7500"
	@echo -e "  $(YELLOW)make build-frontend$(NC)   Build React for production only"
	@echo -e "  $(YELLOW)make build-backend$(NC)    Build Go binary only (needs built frontend)"
	@echo -e "  $(YELLOW)make build-tools$(NC)      Build all utility tools (get-airplanes-data, download-routes, xml2json)"
	@echo -e ""
	@echo -e "$(GREEN)Testing:$(NC)"
	@echo -e "  $(YELLOW)make test-straight-line$(NC)       Instructions for testing with straight-line data (dev mode)"
	@echo -e "  $(YELLOW)make test-straight-line-backend$(NC) Run backend with test data (needs React on :3000)"
	@echo -e "  $(YELLOW)make test-straight-line-prod$(NC)   Build & run production version with test data"
	@echo -e ""
	@echo -e "$(GREEN)Utilities:$(NC)"
	@echo -e "  $(YELLOW)make generate-data$(NC)    Generate airplane snapshot with routes"
	@echo -e "  $(YELLOW)make install$(NC)          Install frontend dependencies (npm install)"
	@echo -e "  $(YELLOW)make check-prereqs$(NC)    Verify required tools are installed"
	@echo -e "  $(YELLOW)make clean$(NC)            Delete build artifacts and executable"
	@echo -e "  $(YELLOW)make clean-data$(NC)       Delete airplane snapshots and downloaded routes"
	@echo -e "  $(YELLOW)make help$(NC)             Show this message"
	@echo -e ""
	@echo -e "$(GREEN)Options:$(NC)"
	@echo -e "  $(YELLOW)KEEP_SNAPSHOT=yes$(NC)     Keep existing airplane snapshot (skip regeneration)"
	@echo -e ""
	@echo -e "$(GREEN)Examples:$(NC)"
	@echo -e "  $(YELLOW)make dev$(NC)                     # Instructions for development"
	@echo -e "  $(YELLOW)make prod$(NC)                    # Build with new airplane snapshot"
	@echo -e "  $(YELLOW)make prod KEEP_SNAPSHOT=yes$(NC)  # Build keeping existing snapshot"
	@echo -e "  $(YELLOW)make prod && make run-prod$(NC)   # Build and run production"
	@echo -e ""

# ============================================================================
# DEVELOPMENT TARGETS
# ============================================================================

# Instruction-only target (no actual commands run)
dev:
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e "$(GREEN)Starting Development Environment$(NC)"
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e ""
	@echo -e "$(BLUE)Open TWO separate terminals and run:$(NC)"
	@echo -e ""
	@echo -e "$(YELLOW)Terminal 1 (React dev server):$(NC)"
	@echo -e "  $$ make dev-frontend"
	@echo -e "  React will be on $(GREEN)http://localhost:$(FRONTEND_PORT)$(NC)"
	@echo -e ""
	@echo -e "$(YELLOW)Terminal 2 (Go API + proxy to React):$(NC)"
	@echo -e "  $$ make dev-backend"
	@echo -e "  API will be on $(GREEN)http://localhost:$(BACKEND_PORT)$(NC)"
	@echo -e ""
	@echo -e "$(BLUE)Then visit: $(GREEN)http://localhost:$(BACKEND_PORT)$(NC)$(NC)"
	@echo -e ""
	@echo -e "$(YELLOW)What happens:$(NC)"
	@echo -e "  - React runs on :3000 (hot reload when you edit frontend)"
	@echo -e "  - Go API runs on :7500 and proxies React requests"
	@echo -e "  - Your backend API routes work normally"
	@echo -e ""

# Start React dev server
dev-frontend: install
	@echo -e "$(GREEN)Starting React dev server on port $(FRONTEND_PORT)...$(NC)"
	cd $(FRONTEND_DIR) && npm start

# Start Go API in development mode (proxies React on port 3000)
dev-backend:
	@if [ ! -f $(ROOT_DIR)/$(DATABASE_DIR)/json/planes_snapshot.json ]; then \
		echo -e "$(YELLOW) No airplane data found, generating...$(NC)"; \
		$(MAKE) generate-data; \
	fi
	@echo -e "$(GREEN)Starting Go API on port $(BACKEND_PORT)...$(NC)"
	@echo -e "$(YELLOW)Note: Make sure React dev server is running on port $(FRONTEND_PORT)$(NC)"
	@echo -e ""
	cd $(ROOT_DIR)/$(BACKEND_DIR) && REACT_URL=http://localhost:$(FRONTEND_PORT) go run ./cmd/api

# ============================================================================
# BUILD TOOLS
# ============================================================================

# Build all utility tools
build-tools:
	@echo -e "$(GREEN)Building utility tools...$(NC)"
	@mkdir -p $(ROOT_DIR)/$(BINARY_DIR)
	@echo -e "$(YELLOW)Building get-airplanes-data...$(NC)"
	@cd $(BACKEND_DIR) && go build -o ../$(GET_AIRPLANES_BIN) ./cmd/get-airplanes-data
	@echo -e "$(GREEN)- Built: $(GET_AIRPLANES_BIN)$(NC)"
	@echo -e "$(YELLOW)Building download-routes...$(NC)"
	@cd $(BACKEND_DIR) && go build -o ../$(DOWNLOAD_ROUTES_BIN) ./cmd/download-routes
	@echo -e "$(GREEN)- Built: $(DOWNLOAD_ROUTES_BIN)$(NC)"
	@echo -e "$(YELLOW)Building xml2json...$(NC)"
	@cd $(BACKEND_DIR) && go build -o ../$(XML2JSON_BIN) ./cmd/xml2json
	@echo -e "$(GREEN)- Built: $(XML2JSON_BIN)$(NC)"
	@echo -e "$(GREEN)- All tools built successfully$(NC)"

# ============================================================================
# PRODUCTION TARGETS
# ============================================================================

# Generate required JSON data files
generate-data: build-tools
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e "$(GREEN)Generating required data files...$(NC)"
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
ifeq ($(KEEP_SNAPSHOT),yes)
	@if [ -f db/json/planes_snapshot.json ]; then \
		echo -e "$(BLUE)ℹ Keeping existing airplane snapshot (KEEP_SNAPSHOT=yes)$(NC)"; \
	else \
		echo -e "$(YELLOW) No existing snapshot found, generating a new one...$(NC)"; \
		echo -e "$(YELLOW)Fetching airplane data and enriching with routes$(NC)"; \
		$(GET_AIRPLANES_BIN) -enrich-routes -output $(ROOT_DIR)/db/json/planes_snapshot.json -routes-dir $(ROOT_DIR)/db/routes; \
		echo -e "$(GREEN)- Airplane snapshot generated and enriched with routes$(NC)"; \
	fi
else
	@echo -e "$(YELLOW)Fetching airplane data and enriching with routes$(NC)"
	@$(GET_AIRPLANES_BIN) -enrich-routes -output $(ROOT_DIR)/db/json/planes_snapshot.json -routes-dir $(ROOT_DIR)/db/routes
	@echo -e "$(GREEN)- Airplane snapshot generated and enriched with routes$(NC)"
endif
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e "$(GREEN)- All data files generated successfully$(NC)"
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"

# Build frontend for production
build-frontend: install
	@echo -e "$(GREEN)Building React app for production...$(NC)"
	cd $(FRONTEND_DIR) && npm run build
	@echo -e "$(GREEN)- Frontend built to: $(FRONTEND_DIR)/build$(NC)"

# Build Go backend (requires frontend to be built first)
build-backend: generate-data build-frontend
	@echo -e "$(GREEN)Building Go binary...$(NC)"
	@mkdir -p $(ROOT_DIR)/$(BINARY_DIR)
	cd $(BACKEND_DIR) && go build -o ../$(BINARY_DIR)/$(BINARY_NAME) ./cmd/api
	@echo -e "$(GREEN)- Binary built: $(BINARY_DIR)/$(BINARY_NAME)$(NC)"

# Build everything for production (single executable)
prod: build-tools build-backend
	@echo -e ""
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e "$(GREEN)Production Build Complete!$(NC)"
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e ""
	@echo -e "$(YELLOW)Executable: $(GREEN)$(BINARY_DIR)/$(BINARY_NAME)$(NC)$(YELLOW) ($(BACKEND_PORT))$(NC)"
	@echo -e ""
	@echo -e "Run with:"
	@echo -e "  $$ make run-prod"
	@echo -e ""

# Run the production executable
run-prod:
	@echo -e "$(GREEN)Starting AI30 AirTraffic on port $(BACKEND_PORT)...$(NC)"
	@echo -e "$(YELLOW)Visit: $(GREEN)http://localhost:$(BACKEND_PORT)$(NC)$(NC)"
	@echo -e ""
	PORT=$(BACKEND_PORT) ./$(BINARY_DIR)/$(BINARY_NAME)

# ============================================================================
# TESTING TARGETS
# ============================================================================

# Run simulation with straight-line test data (in dev mode with React)
test-straight-line: install
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e "$(GREEN)Starting Simulation with Straight-Line Test Data$(NC)"
	@echo -e "$(GREEN)═══════════════════════════════════════════════════════════$(NC)"
	@echo -e ""
	@if [ ! -f $(ROOT_DIR)/$(DATABASE_DIR)/json/straight_line_test_data.json ]; then \
		echo -e "$(RED)✗ Test data file not found: $(DATABASE_DIR)/json/straight_line_test_data.json$(NC)"; \
		exit 1; \
	fi
	@echo -e "$(YELLOW)Test data: $(DATABASE_DIR)/json/straight_line_test_data.json$(NC)"
	@echo -e ""
	@echo -e "$(BLUE)Open TWO separate terminals and run:$(NC)"
	@echo -e ""
	@echo -e "$(YELLOW)Terminal 1 (React dev server):$(NC)"
	@echo -e "  $$ make dev-frontend"
	@echo -e ""
	@echo -e "$(YELLOW)Terminal 2 (Go API with test data):$(NC)"
	@echo -e "  $$ make test-straight-line-backend"
	@echo -e ""
	@echo -e "$(BLUE)Then visit: $(GREEN)http://localhost:$(BACKEND_PORT)$(NC)$(NC)"
	@echo -e ""

# Backend only with test data (requires React dev server running)
test-straight-line-backend:
	@if [ ! -f $(ROOT_DIR)/$(DATABASE_DIR)/json/straight_line_test_data.json ]; then \
		echo -e "$(RED)✗ Test data file not found: $(DATABASE_DIR)/json/straight_line_test_data.json$(NC)"; \
		exit 1; \
	fi
	@echo -e "$(GREEN)Starting Go API with straight-line test data on port $(BACKEND_PORT)...$(NC)"
	@echo -e "$(YELLOW)Note: Make sure React dev server is running on port $(FRONTEND_PORT)$(NC)"
	@echo -e ""
	cd $(ROOT_DIR)/$(BACKEND_DIR) && TEST_DATA_PATH=../db/json/straight_line_test_data.json REACT_URL=http://localhost:$(FRONTEND_PORT) go run ./cmd/api

# Build production version with test data
test-straight-line-prod: build-frontend
	@echo -e "$(GREEN)Building production version with test data...$(NC)"
	@if [ ! -f $(ROOT_DIR)/$(DATABASE_DIR)/json/straight_line_test_data.json ]; then \
		echo -e "$(RED)✗ Test data file not found: $(DATABASE_DIR)/json/straight_line_test_data.json$(NC)"; \
		exit 1; \
	fi
	@mkdir -p $(ROOT_DIR)/$(BINARY_DIR)
	cd $(BACKEND_DIR) && go build -o ../$(BINARY_DIR)/$(BINARY_NAME)-test ./cmd/api
	@echo -e "$(GREEN)- Binary built: $(BINARY_DIR)/$(BINARY_NAME)-test$(NC)"
	@echo -e ""
	@echo -e "$(GREEN)Starting simulation with test data on port $(BACKEND_PORT)...$(NC)"
	@echo -e "$(YELLOW)Visit: $(GREEN)http://localhost:$(BACKEND_PORT)$(NC)$(NC)"
	@echo -e ""
	cd $(ROOT_DIR) && PORT=$(BACKEND_PORT) TEST_DATA_PATH=./$(DATABASE_DIR)/json/straight_line_test_data.json ./$(BINARY_DIR)/$(BINARY_NAME)-test

# ============================================================================
# UTILITY TARGETS
# ============================================================================

# Check prerequisites
check-prereqs:
	@echo -e "$(YELLOW)Checking prerequisites...$(NC)"
	@command -v go >/dev/null 2>&1 || { echo -e "$(RED)✗ Go is not installed. Visit: https://go.dev/dl/$(NC)"; exit 1; }
	@command -v npm >/dev/null 2>&1 || { echo -e "$(RED)✗ npm is not installed. Visit: https://nodejs.org/$(NC)"; exit 1; }
	@command -v node >/dev/null 2>&1 || { echo -e "$(RED)✗ Node.js is not installed. Visit: https://nodejs.org/$(NC)"; exit 1; }
	@echo -e "$(GREEN)- Go: $$(go version | awk '{print $$3}')$(NC)"
	@echo -e "$(GREEN)- Node.js: $$(node --version)$(NC)"
	@echo -e "$(GREEN)- npm: $$(npm --version)$(NC)"
	@echo -e "$(GREEN)- All prerequisites are installed$(NC)"

# Install frontend dependencies
install: check-prereqs
	@if [ ! -d "$(ROOT_DIR)/$(FRONTEND_DIR)/node_modules" ]; then \
		echo -e "$(GREEN)Installing frontend dependencies...$(NC)"; \
		cd $(ROOT_DIR)/$(FRONTEND_DIR) && npm install; \
	else \
		echo -e "$(YELLOW)node_modules already exists, skipping npm install$(NC)"; \
	fi

# Clean build artifacts
clean:
	@echo -e "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -f $(ROOT_DIR)/$(BINARY_DIR)/$(BINARY_NAME)
	@rm -f $(ROOT_DIR)/$(BINARY_DIR)/$(BINARY_NAME)-test
	@rm -rf $(ROOT_DIR)/$(FRONTEND_DIR)/build
	@if [ -d "$(ROOT_DIR)/$(FRONTEND_DIR)/node_modules" ]; then \
		echo -e "$(YELLOW)Removing node_modules (this may take a while)...$(NC)"; \
		rm -rf $(ROOT_DIR)/$(FRONTEND_DIR)/node_modules; \
	fi
	@echo -e "$(GREEN)- Clean complete$(NC)"

# Clean data
clean-data:
	@if [ -d "$(ROOT_DIR)/$(DATABASE_DIR)/json" ]; then \
		echo -e "$(YELLOW)Removing airplane snapshots...$(NC)"; \
		rm -rf $(ROOT_DIR)/$(DATABASE_DIR)/json/planes_snapshot.json; \
	fi
	@if [ -d "$(ROOT_DIR)/$(DATABASE_DIR)/routes" ]; then \
		echo -e "$(YELLOW)Removing downloaded routes...$(NC)"; \
		rm -rf $(ROOT_DIR)/$(DATABASE_DIR)/routes/; \
	fi

# ============================================================================
# DEFAULT TARGET
# ============================================================================
# If user types just 'make' with no arguments, run help
.DEFAULT_GOAL := help