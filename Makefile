.PHONY: build gen-error next-error-code lint test check-sql install-tools migrate-up migrate-down-1 mock-gen

EVENTCORE := $(shell go list -m -f "{{.Dir}}" github.com/QuangTung97/eventcore)

build:
	go build -o bin/errors cmd/errors/main.go
	go build -o bin/migrate cmd/migrate/main.go
	go build -o bin/server cmd/server/main.go
	go build -o bin/event cmd/event/main.go

run-pretty:
	go run cmd/server/main.go start 2>&1 > /dev/null | jq -r ".,.stacktrace"

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

check-sql:
	go run cmd/server/main.go check-sql

install-tools:
	go install github.com/kisielk/errcheck
	go install golang.org/x/lint/golint

migrate-up:
	go run cmd/migrate/main.go up

migrate-down-1:
	go run cmd/migrate/main.go down 1

mock-gen:
	go generate ./...

event-gen:
	genny -in=${EVENTCORE}/core.go -out=todoapp/event/core/core.go -pkg="core" gen "Event=Event"
	genny -in=${EVENTCORE}/options.go -out=todoapp/event/core/options.go -pkg="core" gen "Event=Event"
