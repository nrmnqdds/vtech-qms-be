include .env

lint:
	@echo "Running linter..."
	@golangci-lint run --fix --verbose

goose-create:
	@echo "Creating migration..."
	@goose -dir db/migrations -s create $(NAME) sql

goose/up:
	@echo "Running migrations..."
	@goose -dir db/migrations postgres $(DATABASE_URL) up

goose/down:
	@echo "Running rollback..."
	@goose -dir db/migrations postgres $(DATABASE_URL) down

sqlc/generate:
	@echo "Generating sqlc..."
	@sqlc generate
