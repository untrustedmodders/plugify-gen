#!/bin/bash

echo "Building Plugify Generator for WebAssembly..."

# Create dist directory if it doesn't exist
mkdir -p dist

# Build the WebAssembly binary
GOOS=js GOARCH=wasm go build -o dist/plugify-gen.wasm ./cmd/wasm/main.go

if [ $? -eq 0 ]; then
    echo ""
    echo "Successfully built plugify-gen.wasm"
    echo ""

    # Copy wasm_exec.js from Go installation
    echo "Copying wasm_exec.js..."
    GOROOT_PATH=$(go env GOROOT)

    if [ -f "$GOROOT_PATH/lib/wasm/wasm_exec.js" ]; then
        cp "$GOROOT_PATH/lib/wasm/wasm_exec.js" dist/wasm_exec.js
    elif [ -f "$GOROOT_PATH/misc/wasm/wasm_exec.js" ]; then
        cp "$GOROOT_PATH/misc/wasm/wasm_exec.js" dist/wasm_exec.js
    else
        echo "Error: Could not find wasm_exec.js in Go installation"
        exit 1
    fi

    if [ $? -eq 0 ]; then
        echo ""
        echo "Build complete! Files are in the 'dist' directory:"
        echo "  - plugify-gen.wasm"
        echo "  - wasm_exec.js"
        echo ""
        echo "See WASM.md for integration instructions"
    else
        echo "Error: Could not copy wasm_exec.js"
        echo "Please ensure Go is installed correctly"
        exit 1
    fi
else
    echo "Build failed!"
    exit 1
fi

chmod +x build-wasm.sh