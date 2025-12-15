help: Show this help message
	@echo "Available targets:"
	@echo "  build    Build the BioHub application"
	@echo "  help     Show this help message"


build:
	go build -o biohub cmd/build/main.go && ./biohub

clean:
	rm -f biohub 
	rm -rf dist