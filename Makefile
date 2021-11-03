SHELL=/bin/bash -o pipefail

.PHONY: mockgen
mockgen:
		mockgen -source pkg/meroxa/meroxa.go -imports meroxa=github.com/meroxa/meroxa-go -package mock > pkg/mock/mock_client.go
