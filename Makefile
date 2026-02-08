.PHONY: help build format test cov cov-log vet

BINARY_NAME = biohub

# Nix wrapper logic: Use nix-shell if available and not already inside one
# Also check if we are in a CI environment where we usually want to use system tools
USE_NIX = $(shell if command -v nix-shell >/dev/null 2>&1 && [ -z "$$IN_NIX_SHELL" ] && [ "$$GITHUB_ACTIONS" != "true" ]; then echo "yes"; else echo "no"; fi)

ifeq ($(USE_NIX),yes)
    NIX_RUN = nix-shell --run
else
    NIX_RUN = bash -c
endif

help:
	@echo "Available targets:"
	@echo "  format        Format the Go source code"
	@echo "  vet           Run go vet"
	@echo "  test          Run unit tests"
	@echo "  cov           Generate test coverage report"
	@echo "  cov-log       Generate and display test coverage report"
	@echo "  build         Build the BioHub application"
	@echo "  help          Show this help message"

format:
	@$(NIX_RUN) "go fmt ./cmd/..."

vet:
	@$(NIX_RUN) "go vet ./cmd/..."

test:
	@$(NIX_RUN) "go test ./cmd/... -v"

cov:
	@$(NIX_RUN) "go test -cover ./cmd/..."

cov-log:
	@$(NIX_RUN) "go test -coverprofile=coverage.out ./cmd/... && go tool cover -func=coverage.out && rm coverage.out || exit 1"

build:
	@$(NIX_RUN) "go build -o $(BINARY_NAME) cmd/build/main.go" && ./$(BINARY_NAME)
	@rm -f $(BINARY_NAME)