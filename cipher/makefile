SHELL := /bin/bash
.PHONY: api-run cli-run clean docker-build test tidy fmt

# ==============================================================================
# Building without container
api-run:
	go run api/main.go

cli-run: clean
	go build -o ./cmd/bin/cipher cmd/main.go

clean:
	rm -rf ./cmd/bin

# ==============================================================================
# Building container

VERSION := 1.0

docker-build:
	docker build \
		-f zarf/docker/dockerfile \
		-t cipher-api \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

docker-run:
	docker run -d -p 8080:8080 cipher-api

# ==============================================================================
# Testing coverage support

test:
	go test ./... -v -count=1

# ==============================================================================
# Modules support

tidy:
	go mod tidy


# ==============================================================================
# Code Format support
fmt:
	go fmt ./...

