build:
	@go build -o bin/paclarative
run: build
	@./bin/paclarative
test:
	@go test -v ./...
