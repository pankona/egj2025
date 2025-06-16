.PHONY: lint test

lint:
	go vet ./...

test:
	GOOS=windows go test -v ./...