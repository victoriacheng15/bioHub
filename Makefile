help: Show this help message
	@echo "Available targets:"
	@echo "  build    Build the BioHub application"
	@echo "  clean    Clean up build artifacts"
	@echo "  format   Format the Go source code"
	@echo "  help     Show this help message"


build:
	go build -o biohub cmd/build/main.go && ./biohub

clean:
	rm -f biohub 
	rm -rf dist

format:
	go fmt -w ./cmd