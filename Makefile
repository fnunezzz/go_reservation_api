build:
	@go build -o bin/go_reservation_api

run: export DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres
run: export MIGRATIONS_PATH=./sql
run: build
	@./bin/go_reservation_api

test:
	@go test -v ./...