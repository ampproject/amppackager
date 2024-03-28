.PHONY: default clean check test

default: clean check test build

test: clean
	go test -v -cover ./...

clean:
	rm -f cover.out

build:
	go build

check:
	golangci-lint run
