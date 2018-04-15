.PHONY: compile test clean run

all: compile

compile:
	go build

test:
	go test ./...

run:
	go run main.go

clean:
	go clean
