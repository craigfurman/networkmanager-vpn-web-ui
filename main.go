package main

import (
	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/urfave/negroni"
)

func main() {
	api := vpnmanager.NewAPI(&vpnmanager.NmcliClient{})
	server := negroni.Classic()
	server.UseHandler(api)
	server.Run(":8080") // TODO parameterise bind port/address
}
