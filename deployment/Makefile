# Define variables
ENV_FILE = ../.env

# Export .env to environment
init:
	@echo "Exporting environment variables from $(ENV_FILE)..."
	@export $(shell sed 's/=.*//' $(ENV_FILE))

# Run goose up command
up:
	@echo "Running goose up..."
	@goose up

# Run goose down command
down:
	@echo "Running goose down..."
	@goose down

# Build cmd/go-uber-fx/main.go
build:
	@echo "Building cmd/go-uber-fx/main.go..."
	@go build -o main cmd/go-uber-fx/main.go

# Run cmd/go-uber-fx/main.go
run:
	@echo "Running cmd/go-uber-fx/main.go..."
	@go run cmd/go-uber-fx/main.go
