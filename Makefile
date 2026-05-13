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

emit-audit-event:
	curl -X POST http://localhost:8080/events \
      -H "Content-Type: application/json" \
      -d '{"id":"1","source":"auth","action":"login","actor":"alex","timestamp":"2026-05-13T10:00:00Z"}'
