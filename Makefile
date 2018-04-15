.PHONY: compile test clean

all: compile

compile:
	go build

test:
	go test ./...

clean:
	go clean
