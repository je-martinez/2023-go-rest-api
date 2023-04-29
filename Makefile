# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./app

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go


# ==============================================================================
# Main

build:
	go build main .
run: build
	go run main

test:
	go test -cover ./...

# ==============================================================================
# Docker

up:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d

down:
	echo "Removing local environment"
	docker-compose -f docker-compose.local.yml down -v

logs:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml logs -f