package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/urfave/negroni"
)

// Overridden at link time on release, cannot non-string
var dist = "false"

func main() {
	bindAddress := flag.String("bind-address", ":8080", "bind address")
	flag.Parse()

	staticFileDir := "public"
	if dist == "true" {
		staticFileDir = filepath.Join(filepath.Dir(os.Args[0]), "public")
	}

	api := vpnmanager.NewAPI(&vpnmanager.NmcliClient{})
	server := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(staticFileDir)),
	)
	server.UseHandler(api)
	server.Run(*bindAddress)
}
