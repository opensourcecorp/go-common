SHELL = /usr/bin/env bash -euo pipefail

all: test

.PHONY: %

test:
	@go test -v -cover ./...
