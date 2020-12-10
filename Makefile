.PHONY: build gen-error next-error-code lint test

build:
	go build -o bin/errors cmd/errors/main.go

gen-error:
	go run cmd/errors/main.go generate

next-error-code:
	go run cmd/errors/main.go next-code $(rpc-status)

lint:
	go fmt ./...
	golint ./...
	go vet ./...
	errcheck ./...

test:
	go test -v ./...
