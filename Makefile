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

install-tools:
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install golang.org/x/tools/cmd/goimports@latest

fmt: install-tools
	yamlfmt .
	goimports -w .

clean:
	rm -rf dist
