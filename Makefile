include .env

lint:
	@echo "Running linter..."
	@golangci-lint run --fix --verbose

goose-create:
	@echo "Creating migration..."
	@goose -dir db/migrations -s create $(NAME) sql

migrate:
	@echo "Running migrations..."
	@goose -dir db/migrations postgres $(DATABASE_URL) up

rollback:
	@echo "Running rollback..."
	@goose -dir db/migrations postgres $(DATABASE_URL) down

gen:
	@echo "Generating sqlc..."
	@sqlc generate

test:
	@echo "Running tests..."
	@go test -v -cover ./...

mock:
	@echo "Generating mocks..."
	@mockery

# go install github.com/dkorunic/betteralign/cmd/betteralign@latest
align:
	@echo "Aligning field alignments..."
	betteralign -apply ./...

