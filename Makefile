help: Show this help message
	@echo "Available targets:"
	@echo "  build         Build the BioHub application"
	@echo "  clean         Clean up build artifacts"
	@echo "  format        Format the Go source code"
	@echo "  test          Run unit tests"
	@echo "  coverage      Generate test coverage report"
	@echo "  coverage-html Generate HTML test coverage report"
	@echo "  help          Show this help message"


build:
	go build -o biohub.exe cmd/build/main.go && ./biohub.exe && rm ./biohub.exe

clean:
	rm -f biohub.exe 
	rm -rf dist

format:
	go fmt ./cmd/...

test:
	go test ./cmd/... -v

coverage:
	go test -cover ./cmd/...

coverage-html:
	go test -coverprofile=coverage.out ./cmd/... && go tool cover -html=coverage.out