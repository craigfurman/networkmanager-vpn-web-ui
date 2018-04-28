package main

import (
	"flag"

	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/urfave/negroni"
)

func main() {
	bindAddress := flag.String("bind-address", ":8080", "bind address")
	flag.Parse()

	api := vpnmanager.NewAPI(&vpnmanager.NmcliClient{})
	server := negroni.Classic()
	server.UseHandler(api)
	server.Run(*bindAddress)
}
