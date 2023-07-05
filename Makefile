.PHONY: test
test:
	@go test -v -race ./...

.PHONY: test_kvs
test_kvs:
	@go test -v ./cmd/...

kvs: ./cmd/kvs/*.go
	@CGO_ENABLED=0 GOOS=linux go build -o kvs ./cmd/kvs