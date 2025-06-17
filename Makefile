.PHONY: lint test build-wasm serve-wasm clean

lint:
	go vet ./...

test:
	go test -v ./...

build-wasm:
	mkdir -p dist
	GOOS=js GOARCH=wasm go build -o dist/main.wasm .
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js dist/
	cp web/index.html dist/

serve-wasm: build-wasm
	go run github.com/hajimehoshi/wasmserve@latest -http=:8080 .

clean:
	rm -rf dist
