#!/bin/bash
gofmt -d -s *.go
docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run
