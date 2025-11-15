@echo off
echo Building Plugify Generator for WebAssembly...

REM Set GOOS and GOARCH for WebAssembly
set GOOS=js
set GOARCH=wasm

REM Build the WebAssembly binary
go build -o dist\plugify-gen.wasm .\cmd\wasm\main.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Successfully built plugify-gen.wasm
    echo.

    REM Copy wasm_exec.js from Go installation
    echo Copying wasm_exec.js...
    if not exist dist mkdir dist
    if exist "%GOROOT%\lib\wasm\wasm_exec.js" (
        copy "%GOROOT%\lib\wasm\wasm_exec.js" dist\wasm_exec.js
    ) else if exist "%GOROOT%\misc\wasm\wasm_exec.js" (
        copy "%GOROOT%\misc\wasm\wasm_exec.js" dist\wasm_exec.js
    ) else (
        echo Error: Could not find wasm_exec.js in Go installation
        exit /b 1
    )

    if %ERRORLEVEL% EQU 0 (
        echo.
        echo Build complete! Files are in the 'dist' directory:
        echo   - plugify-gen.wasm
        echo   - wasm_exec.js
        echo.
        echo See WASM.md for integration instructions
    ) else (
        echo Error: Could not copy wasm_exec.js
        echo Please ensure GOROOT is set correctly
    )
) else (
    echo Build failed!
    exit /b 1
)