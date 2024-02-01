# Makefile

.PHONY: all build run test clean

BINARY_NAME=movierental

all: build

build:
	go build -o bin/$(BINARY_NAME) cmd/main.go

run:
	bin/$(BINARY_NAME)

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

test:
	go clean -testcache && 	go test ./...