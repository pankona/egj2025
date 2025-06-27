.PHONY: lint test build-wasm serve-wasm clean generate-stages

lint:
	GOOS=js GOARCH=wasm go vet ./...

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

generate-stages:
	@echo "Generating stage files from ASCII art..."
	@for i in 01 02 03 04 05 06 07 08 09 10; do \
		if [ -f stage$$i.txt ]; then \
			echo "Generating stage$$i.go from stage$$i.txt"; \
			go run cmd/stagegen/main.go stage$$i.txt; \
			if [ -f stage$$i.go ]; then \
				target_name=stage$$(echo $$i | sed 's/^0*//').go; \
				if [ "stage$$i.go" != "$$target_name" ]; then \
					mv stage$$i.go $$target_name; \
				fi; \
			fi; \
		else \
			echo "Warning: stage$$i.txt not found"; \
		fi; \
	done
	@echo "Stage generation complete"

clean:
	rm -rf dist
