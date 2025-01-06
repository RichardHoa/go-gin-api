test:
	@go test -v ./...

get:
	@go get ./...

build:
	@go mod tidy
	@go build -o go-api cmd/main.go

migrate:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrations/main.go up

migrate-down:
	@go run cmd/migrate/migrations/main.go down