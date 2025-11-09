build:
	@go build -o ./bin/sh

run: build
	@./bin/sh

test: 
	@go test ./... -v