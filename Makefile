.PHONY: build generate fmt lint tidy test

build:
	go build -v ./...

deps:
	go install ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

fmt:
	find -iname '*.go' | xargs -L1 gofmt -s -w

tidy:
	go mod tidy

lint:
	golangci-lint run

test:
	ginkgo -r -race -randomizeAllSpecs -randomizeSuites -trace -progress -cover -skipPackage ./...