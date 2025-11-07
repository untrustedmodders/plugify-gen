.PHONY: build test clean install help

BINARY_NAME=plugify-gen
INSTALL_PATH=/usr/local/bin

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the generator binary
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/plugify-gen
	@echo "✓ Built $(BINARY_NAME)"

test: build ## Test the generator with example manifest
	@echo "Testing C++ generator..."
	@./$(BINARY_NAME) -manifest plugify-plugin-s2sdk.pplugin.in -output ./test_output/cpp -lang cpp -overwrite -verbose
	@echo ""
	@echo "Testing V8 generator..."
	@./$(BINARY_NAME) -manifest plugify-plugin-s2sdk.pplugin.in -output ./test_output/v8 -lang v8 -overwrite -verbose
	@echo ""
	@echo "✓ All tests passed"

clean: ## Clean build artifacts and test outputs
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf test_output/
	@echo "✓ Cleaned"

install: build ## Install the binary to system path
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo cp $(BINARY_NAME) $(INSTALL_PATH)/
	@echo "✓ Installed to $(INSTALL_PATH)/$(BINARY_NAME)"

uninstall: ## Uninstall the binary from system path
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "✓ Uninstalled"

benchmark: build ## Benchmark against Python generators
	@echo "Benchmarking C++ generator..."
	@echo "Go implementation:"
	@time ./$(BINARY_NAME) -manifest plugify-plugin-s2sdk.pplugin.in -output ./test_output/cpp -lang cpp -overwrite > /dev/null
	@echo ""
	@echo "Python implementation:"
	@time python generator._cpppy plugify-plugin-s2sdk.pplugin.in ./test_output/python > /dev/null 2>&1 || true
	@echo ""
	@echo "✓ Benchmark complete"

# Development targets
fmt: ## Format Go code
	@go fmt ./...

vet: ## Run go vet
	@go vet ./...

lint: fmt vet ## Run all linters

dev-test: lint test ## Run linters and tests

# Quick shortcuts
run-cpp: build ## Quick test: generate C++ bindings
	@./$(BINARY_NAME) -manifest plugify-plugin-s2sdk.pplugin.in -output ./out -lang cpp -overwrite

run-v8: build ## Quick test: generate V8 bindings
	@./$(BINARY_NAME) -manifest plugify-plugin-s2sdk.pplugin.in -output ./out -lang v8 -overwrite
