SHELL           =/bin/bash -o pipefail
MOCKGEN_VERSION ?= v1.6.0

mockgen-install:
	go install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

.PHONY: mockgen
mockgen: mockgen-install
		mockgen -source pkg/meroxa/meroxa.go -imports meroxa=github.com/meroxa/meroxa-go -package mock > pkg/mock/mock_client.go

.PHONY: gomod
gomod:
	GOPRIVATE=github.com/meroxa/merman,github.com/meroxa/catalyst,github.com/meroxa/piper,github.com/meroxa/x go mod tidy
	go mod vendor

.PHONY: test
test: mockgen
	go test ./...

.PHONY: lint
lint: ./bin/golangci-lint
	./bin/golangci-lint run

.PHONY: ci-lint
ci-lint: ./bin/golangci-lint
	./bin/golangci-lint run --out-format=github-actions -v

./bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.49.0

.PHONY: vet
vet: lint