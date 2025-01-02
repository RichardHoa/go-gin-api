test:
	@go test -v ./...

i:
	@go get ./...

migrate:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrations/main.go up

migrate-down:
	@go run cmd/migrate/migrations/main.go down