//go:build js && wasm

package main

import (
	"syscall/js"
)

// initDebugMode initializes debug mode based on URL parameters for WASM builds
func initDebugMode() {
	// Get window.location.search
	window := js.Global().Get("window")
	location := window.Get("location")
	search := location.Get("search").String()

	// Parse URL parameters
	urlParams := js.Global().Get("URLSearchParams").New(search)
	debugParam := urlParams.Call("get", "debug").String()

	// Set debug mode based on parameter
	DebugMode = debugParam == "true" || debugParam == "1"
}
