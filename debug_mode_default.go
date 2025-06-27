//go:build !js || !wasm

package main

import (
	"os"
)

// initDebugMode initializes debug mode based on environment variables for non-WASM builds
func initDebugMode() {
	// Check DEBUG environment variable
	debugEnv := os.Getenv("DEBUG")
	DebugMode = debugEnv == "true" || debugEnv == "1"
}
