VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')
DOCKER := $(shell which docker)

export GO111MODULE = on

all: ci-lint ci-test install

###############################################################################
# Build / Install
###############################################################################

LD_FLAGS = -X github.com/forbole/juno/v2/cmd.Version=$(VERSION) \
	-X github.com/forbole/juno/v2/cmd.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building wasmx binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/wasmx.exe ./cmd/wasmx
else
	@echo "building wasmx binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/wasmx ./cmd/wasmx
endif

install: go.sum
	@echo "installing wasmx binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/wasmx

###############################################################################
# Tests / CI
###############################################################################

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" | xargs goimports -w -local github.com/disperze/wasmx
.PHONY: format

clean:
	rm -f tools-stamp ./build/**

.PHONY: install build ci-test ci-lint clean
