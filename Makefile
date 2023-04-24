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

go-build:
	go build ./cmd/api/main.go

go-run: go-build
	go run ./cmd/api/main.go

go-test:
	go test -cover ./...

# ==============================================================================
# Docker

run:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d

stop:
	echo "Removing local environment"
	docker-compose -f docker-compose.local.yml down -v

logs:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml logs -f