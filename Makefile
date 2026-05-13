.PHONY: help build test lint

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build   - Build the project"
	@echo "  test    - Run tests"
	@echo "  lint    - Run linters"

build:
	go build ./...

test:
	go test -v ./...

lint:
	golangci-lint run -v
