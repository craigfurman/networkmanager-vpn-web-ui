.PHONY: compile test clean run vendordeps

all: compile

compile:
	go build

test:
	go test ./...

run:
	go run main.go

clean:
	go clean

vendordeps:
	govendor add +external
