TEST 		?= ./...

ROOT		:= $(shell cd ../../../.. && pwd)
GOPATH 	:= $(GOPATH):$(ROOT)

default: test

fmt: generate
	go fmt ./...

test: generate
	go test $(TEST) $(TESTARGS)

generate:
	go generate ./...

updatedeps:
	go get -u golang.org/x/tools/cmd/stringer

.PHONY: default generate test updatedeps
