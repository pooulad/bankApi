build:
	@go build -o ./bin/bankApi

run: build
	@./bin/bankApi

test:
	@go test -v ./...