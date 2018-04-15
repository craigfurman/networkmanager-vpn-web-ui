package vpnmanager

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPI(netManager NetworkManager) *VpnConnectionManager {
	manager := &VpnConnectionManager{
		NetworkManager: netManager,
		Router:         mux.NewRouter(),
	}
	manager.HandleFunc("/api/connections", manager.listConnections).Methods(http.MethodGet)

	return manager
}

type NetworkManager interface {
	ListConnections() ([]Connection, error)
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
	if err != nil {
		// TODO test
		panic(err)
	}

	respW.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(respW).Encode(conns); err != nil {
		// TODO don't panic?
		panic(err)
	}
}
