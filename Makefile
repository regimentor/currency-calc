

server dev:
	go run ./cmd/currency-calc

install:
	go mod download

run-test:
	go test -coverprofile=coverage.out ./...

build compile:
	rm -rf ./build && go build -o ./build/main ./cmd/currency-calc

build-linux64 compile-linux64:
	rm -rf ./build && GOOS=linux GOARCH=amd64 go build -o ./build/main ./cmd/currency-calc