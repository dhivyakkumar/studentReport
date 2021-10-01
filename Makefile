.PHONY: build
build:
	@go build ./...

.PHONY: test
test:
	@go test ./src/controller/ -v