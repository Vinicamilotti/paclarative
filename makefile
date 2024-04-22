build:
	@go build -o bin/server
run: build
	@./bin/paclarative
test:
	@go test -v ./...
