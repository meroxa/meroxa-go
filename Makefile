SHELL=/bin/bash -o pipefail

.PHONY: mockgen
mockgen:
		mockgen -source pkg/meroxa/meroxa.go -imports meroxa=github.com/meroxa/meroxa-go -package mock > pkg/mock/mock_client.go

.PHONY: gomod
gomod:
	GOPRIVATE=github.com/meroxa/merman,github.com/meroxa/catalyst,github.com/meroxa/piper,github.com/meroxa/x go mod tidy
	go mod vendor

.PHONY: test
test:
	go test ./...
