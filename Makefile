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
version = 0.1.0

dist: compile
	mkdir $(name)-$(version)
	cp networkmanager-vpn-web-ui $(name)-$(version)
	cp -a public $(name)-$(version)
	tar -czf $(name)-$(version).tar.gz networkmanager-vpn-web-ui-$(version)
	rm -rf $(name)-$(version)
