.PHONY: help nix-% build clean format test cov cov-html

GO ?= go
BINARY_NAME = biohub

help:
	@echo "Available targets:"
	@echo "  nix-%         Run a command within the nix-shell"
	@echo "  build         Build the BioHub application"
	@echo "  clean         Clean up build artifacts"
	@echo "  format        Format the Go source code"
	@echo "  test          Run unit tests"
	@echo "  cov           Generate test coverage report"
	@echo "  cov-html      Generate HTML test coverage report"
	@echo "  help          Show this help message"

nix-%:
	@nix-shell --run "make $*"

build:
	@$(GO) build -o $(BINARY_NAME) cmd/build/main.go
	@./$(BINARY_NAME)
	@rm -f $(BINARY_NAME)

clean:
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out
	@rm -rf dist

format:
	@$(GO) fmt ./cmd/...

test:
	@$(GO) test ./cmd/... -v

cov:
	@$(GO) test -cover ./cmd/...

cov-html:
	@$(GO) test -coverprofile=coverage.out ./cmd/... && $(GO) tool cover -html=coverage.out
	@rm -f coverage.out