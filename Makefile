.PHONY: all fmt test

all: build

test:
	go test $$(go list ./... | grep -v /vendor/) -race -coverprofile cover.out -timeout 120s

build:
	go build -o bin/mintless-cli ./cmd/
