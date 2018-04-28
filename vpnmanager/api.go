package vpnmanager

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewAPI(netManager NetworkManager) *VpnConnectionManager {
	manager := &VpnConnectionManager{
		NetworkManager: netManager,
		Router:         mux.NewRouter(),
	}

	manager.HandleFunc("/api/connections", manager.listConnections).Methods(http.MethodGet)
	manager.HandleFunc("/api/connections/{name}", manager.setConnection).Methods(http.MethodPut)

	return manager
}

type NetworkManager interface {
	ListConnections() ([]Connection, error)
	SetConnection(conn Connection) error
}

type Connection struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type VpnConnectionManager struct {
	NetworkManager NetworkManager
	*mux.Router
}

func (c *VpnConnectionManager) listConnections(respW http.ResponseWriter, req *http.Request) {
	conns, err := c.NetworkManager.ListConnections()
	must(err)

	respW.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(respW).Encode(conns))
}

func (c *VpnConnectionManager) setConnection(respW http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	active, err := strconv.ParseBool(req.FormValue("active"))
	must(err)

	must(c.NetworkManager.SetConnection(Connection{Name: name, Active: active}))
}

// Ridiculously naive error handling
func must(err error) {
	if err != nil {
		panic(err)
	}
}
