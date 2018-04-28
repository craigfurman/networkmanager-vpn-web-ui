package vpnmanager

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	ipify "github.com/rdegges/go-ipify"
)

func NewAPI(netManager NetworkManager) *VpnConnectionManager {
	manager := &VpnConnectionManager{
		NetworkManager: netManager,
		Router:         mux.NewRouter(),
	}

	manager.HandleFunc("/api/connections", manager.listConnections).Methods(http.MethodGet)
	manager.HandleFunc("/api/connections/{name}", manager.setConnection).Methods(http.MethodPut)
	manager.HandleFunc("/api/address", getAddress).Methods(http.MethodGet)

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

func getAddress(respW http.ResponseWriter, req *http.Request) {
	ip := getPublicIP()
	respBytes, err := json.Marshal(map[string]string{"ip": ip})
	must(err)
	_, err = respW.Write(respBytes)
	must(err)
}

// As the VPN connection is (dis)established, ongoing TCP sessions (such as the
// one with ipify.org) can hang. Retry promptly.
func getPublicIP() string {
	ipCh := make(chan string)
	go func(ipCh chan<- string) {
		ip, err := ipify.GetIp()
		must(err)
		ipCh <- ip
	}(ipCh)

	select {
	case ip := <-ipCh:
		return ip
	case <-time.After(time.Second):
		return getPublicIP()
	}
}

// Ridiculously naive error handling
func must(err error) {
	if err != nil {
		panic(err)
	}
}
