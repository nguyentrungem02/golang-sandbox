# Run server
server:
	go run ./cmd/api

# Build app prod
build:
	go build -o bin/myapp ./cmd/api

# Run
run-binary:
	./bin/myapp

.PHONY: server build run-binary
