.PHONY: build-linux build-win  run-linux run-win  migrate-create migrate-up migrate-down migrate-force migrate-ver

# Directory for migrations
dir="./cmd/migration/"

# Default target - run the application
dev:
	go run ./cmd/echoes/*.go

# Run the application
run-linux:
	./bin/echoes

run-win:
	./bin/echoes.exe


# Build the Go application
build-linux:
	go build -o bin/echoes ./cmd/echoes/*.go

build-win:
	go build -o bin/echoes.exe ./cmd/echoes/*.go


# Create a new migration
migrate-create:
	@read -p "Enter migration name: " name && migrate create -ext sql -dir $(dir) -seq $$name


# Run migrations up (shortcut)
migrate-up:
	@migrate -database "$$DB_ADDR" -path $(dir) -verbose up $(or $(n),)

# Run migrations down (rollback)
migrate-down:
	@migrate -database "$$DB_ADDR" -path $(dir) -verbose down $(or $(n),)

migrate-force:
	@migrate -database "$$DB_ADDR" -path $(dir) force $(v)

migrate-ver:
	@migrate -database "$$DB_ADDR" -path $(dir) version