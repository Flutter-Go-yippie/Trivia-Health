# Makefile for NeuroCoach AI Project (MongoDB)

# Configuration
PROJECT_NAME := neurocoach
MONGO_USER := user
MONGO_PASSWORD := password
MONGO_DB := fitness_ai
MONGO_HOST := localhost
MONGO_PORT := 27017
MONGO_CONTAINER_NAME := neurocoach-mongo

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: mongo-up
mongo-up: ## Start MongoDB using Docker
	docker run -d --name $(MONGO_CONTAINER_NAME) -p $(MONGO_PORT):27017 -e MONGO_INITDB_ROOT_USERNAME=$(MONGO_USER) -e MONGO_INITDB_ROOT_PASSWORD=$(MONGO_PASSWORD) -e MONGO_INITDB_DATABASE=$(MONGO_DB) mongo:6.0

.PHONY: mongo-down
mongo-down: ## Stop and remove MongoDB container
	docker stop $(MONGO_CONTAINER_NAME) || true
	docker rm $(MONGO_CONTAINER_NAME) || true

.PHONY: mongo-shell
mongo-shell: ## Open a shell to the running MongoDB container
	docker exec -it $(MONGO_CONTAINER_NAME) mongosh -u $(MONGO_USER) -p $(MONGO_PASSWORD) --authenticationDatabase admin $(MONGO_DB)

.PHONY: run-api
run-api: ## Run the Go API server
	@echo "Starting API server..."
	@go run cmd/server/main.go

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: build
build: ## Build the application binary
	@echo "Building application..."
	@go build -o bin/neurocoach cmd/server/main.go

.PHONY: clean
clean: mongo-down ## Clean up resources
	@echo "Cleaning up..."
	@rm -rf bin/

.PHONY: full-start
full-start: mongo-up run-api ## Full start: MongoDB and API
	@echo "System started successfully!"

.PHONY: restart-mongo
restart-mongo: mongo-down mongo-up ## Restart MongoDB
