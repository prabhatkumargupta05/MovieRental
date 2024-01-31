# Makefile

.PHONY: all build run clean

BINARY_NAME=movierental

all: build

build:
	go build -o $(BINARY_NAME) cmd/main.go

run:
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)