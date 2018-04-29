.PHONY: compile test testdocker clean run vendordeps

all: compile

compile:
	go build -ldflags "-X main.dist=true"

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

name = networkmanager-vpn-web-ui

dist: compile
	mkdir -p dist/$(name)
	cp networkmanager-vpn-web-ui LICENSE dist/$(name)
	cp -a public dist/$(name)
	tar --owner root --group root -czf $(name).tar.gz -C dist networkmanager-vpn-web-ui
	rm -rf dist
