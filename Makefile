build:
	@go build -o ./bin/bankApi ./cmd/app/main.go

run: build
	@./bin/bankApi

test:
	@go test -v ./...

tidy:
	@go mod tidy