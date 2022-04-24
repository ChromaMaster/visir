.PHONY: build generate fmt lint tidy test clean

build:
	go build -v ./...

deps:
	go install ./...
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.1.3
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

fmt:
	find -iname '*.go' | xargs -L1 gofmt -s -w

tidy:
	go mod tidy

lint:
	golangci-lint run

test:
	ginkgo ./...

clean:
	find -name "*.coverprofile" -type f | xargs rm
