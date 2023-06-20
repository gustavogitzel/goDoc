ALL_SRC := $(shell find . -type f -name '*.go')

.PHONY: all clean check lint test

all: clean check lint test

check:
	go get -u github.com/kisielk/errcheck
	errcheck ./...

lint:
	go get -u golang.org/x/lint/golint
	golint ./...

clean:
	go mod tidy

test:
	go test ./...
