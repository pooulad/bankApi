build:
	@go build -o ./bin/bankApi ./cmd/app/main.go

run: build
	@./bin/bankApi

tidy:
	@go mod tidy

testgo:
	@go test ./... -v