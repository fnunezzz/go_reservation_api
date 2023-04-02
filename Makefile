build:
	@go build -o bin/go_reservation_api

run: build
	@./bin/go_reservation_api

test:
	@go test -v ./...