# networkmanager-vpn-web-ui

[![CircleCI](https://circleci.com/gh/craigfurman/networkmanager-vpn-web-ui.svg?style=svg)](https://circleci.com/gh/craigfurman/networkmanager-vpn-web-ui)

A web UI for remotely managing NetworkManager-managed VPN connections.

## Why?

My target users are people who want to toggle the VPN connections of a remote
Linux machine that runs NetworkManager, but aren't comfortable with the
terminal. I realise the intersection of that particular Venn diagram is rather
small.

I looked at
[mk-fg/NetworkManager-WiFi-WebUI](https://github.com/mk-fg/NetworkManager-WiFi-WebUI),
a similar project for managing WiFi connections over a web UI. It doesn't manage
VPN connections, and I wanted the simplified deployment of Go over Python, so I
knocked this together instead of contributing.

You should **never** expose this server on any untrusted network! It is not even
slightly secure. It has no authentication/authorization concepts, and the API
performs no validation on inputs.

## Installation

Note that to actually manage connections, NetworkManager must have already
stored the VPN credentials for the user you run the service as.

### From source

`go get github.com/craigfurman/networkmanager-vpn-web-ui`, or otherwise clone
this project into your GOPATH.

To run in dev mode: `make run`.

To create a distributable archive, `make dist`. Untar the resulting tarball
somewhere (e.g. `/opt`), and execute the binary.

### systemd service

Coming soon.

### Arch Linux

Coming soon.

## Contributing

If this project is useful to you, please feel free to open an issue or pull
request. I knocked this together in a few hours, but in the unlikely event that
this is useful to others I'm open to cleaning it up.
