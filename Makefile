.PHONY: help build format test cov cov-html vet

BINARY_NAME = biohub

help:
	@echo "Available targets:"
	@echo "  format        Format the Go source code"
	@echo "  vet           Run go vet"
	@echo "  test          Run unit tests"
	@echo "  cov           Generate test coverage report"
	@echo "  cov-html      Generate HTML test coverage report"
	@echo "  build         Build the BioHub application"
	@echo "  help          Show this help message"

format:
	@nix-shell --run "go fmt ./cmd/..."

vet:
	@nix-shell --run "go vet ./cmd/..."

test:
	@nix-shell --run "go test ./cmd/... -v"

cov:
	@nix-shell --run "go test -cover ./cmd/..."

cov-html:
	@nix-shell --run "go test -coverprofile=coverage.out ./cmd/... && go tool cover -html=coverage.out"
	@rm -f coverage.out

build:
	@nix-shell --run "go build -o $(BINARY_NAME) cmd/build/main.go" && ./$(BINARY_NAME)
	@rm -f $(BINARY_NAME)