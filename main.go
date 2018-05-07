package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/urfave/negroni"
)

// Overridden at link time on release
var staticFilesDir = ""

func main() {
	bindAddress := flag.String("bind-address", ":8080", "bind address")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if staticFilesDir == "" {
		staticFilesDir = filepath.Join(cwd, "public")
	}
	if staticFilesDir == "_" {
		staticFilesDir = filepath.Join(filepath.Dir(os.Args[0]), "public")
	}

	api := vpnmanager.NewAPI(&vpnmanager.NmcliClient{})
	server := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(staticFilesDir)),
	)
	server.UseHandler(api)
	server.Run(*bindAddress)
}
