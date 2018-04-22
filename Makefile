.PHONY: compile test testdocker clean run vendordeps

all: compile

compile:
	go build

test:
	go test ./...

docker_workdir = /go/src/github.com/craigfurman/networkmanager-vpn-web-ui

testdocker:
	docker run --rm -v ${PWD}:$(docker_workdir) -w $(docker_workdir) \
		circleci/golang:1.10.1 make test

run:
	go run main.go

clean:
	go clean

vendordeps:
	govendor add +external
