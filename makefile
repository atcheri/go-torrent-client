.PHONY: build run cli lint test

dependencies:
	go mod download

build:
	go build \
		-o ./out/app \
		cmd/http/main.go

run:
	go run cmd/http/main.go

cli:
	go run cmd/cli/main.go

lint:
	golangci-lint run --issues-exit-code 0 --out-format code-climate | jq -c '.[] | select(.severity|contains("major"))'

test:
	go test -v ./internal/...
