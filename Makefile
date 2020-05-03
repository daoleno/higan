PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=HIGAN \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=higand \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=higancli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

build-linux: go.sum
	env GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/higand ./cmd/higand
	env GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/higancli ./cmd/higancli

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/higand
		go install $(BUILD_FLAGS) ./cmd/higancli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test $(PACKAGES)