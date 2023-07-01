.PHONY: test
test:
	@go test -v -race ./...

.PHONY: test_kvs
test_kvs:
	@go test -v ./cmd/...