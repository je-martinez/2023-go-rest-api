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

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Docker
local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

local-logs:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml logs -f