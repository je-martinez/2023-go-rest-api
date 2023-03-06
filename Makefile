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

go-run:
	go run ./cmd/api/main.go

go-build:
	go build ./cmd/api/main.go

go-test:
	go test -cover ./...

# ==============================================================================
# Docker

run:
	echo "Starting local environment"
	docker-compose up --build -d

stop:
	echo "Removing local environment"
	docker-compose down

logs:
	echo "Starting local environment"
	docker-compose logs -f