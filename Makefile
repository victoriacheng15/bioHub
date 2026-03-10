.PHONY: help build format test test-cov test-cov-log vet update	

BINARY_NAME = biohub

help:
	@echo "Available targets:"
	@echo "  format        Format the Go source code"
	@echo "  vet           Run go vet"
	@echo "  test          Run unit tests"
	@echo "  test-cov      Generate test coverage report"
	@echo "  test-cov-log  Generate and display test coverage report"
	@echo "  build         Build the BioHub application"
	@echo "  update        Update Go dependencies"
	@echo "  help          Show this help message"

format:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v

test-cov:
	go test -cover ./...

test-cov-log:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out && rm coverage.out || exit 1

update:
	go get -u ./... && go mod tidy

build:
	go build -o $(BINARY_NAME) cmd/web/main.go && ./$(BINARY_NAME)
	@rm -f $(BINARY_NAME)