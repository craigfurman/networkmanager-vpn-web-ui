package vpnmanager_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListVPNConnections(t *testing.T) {
	networkManager := &mockNetworkManager{}
	expectedConns := []vpnmanager.Connection{
		{Name: "c1", Active: false},
		{Name: "c2", Active: true},
	}
	networkManager.On("ListConnections").Return(expectedConns, nil)

	server := newTestServer(networkManager)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/connections", server.URL))
	require.Nil(t, err)
	defer resp.Body.Close()
	require.Equal(t, resp.StatusCode, http.StatusOK)

	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var conns []vpnmanager.Connection
	err = json.NewDecoder(resp.Body).Decode(&conns)
	assert.Nil(t, err)
	assert.Equal(t, expectedConns, conns)
}

func TestSetConnection(t *testing.T) {
	networkManager := &mockNetworkManager{}
	expectedConn := vpnmanager.Connection{Name: "a-conn", Active: true}
	networkManager.On("SetConnection", expectedConn).Return(nil)

	server := newTestServer(networkManager)
	defer server.Close()

	uri := fmt.Sprintf("%s/api/connections/a-conn?active=true", server.URL)
	req, err := http.NewRequest(http.MethodPut, uri, nil)
	require.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, http.StatusOK)

	networkManager.AssertCalled(t, "SetConnection", expectedConn)
}

func newTestServer(networkManager vpnmanager.NetworkManager) *httptest.Server {
	httpHandler := vpnmanager.NewAPI(networkManager)
	return httptest.NewServer(httpHandler)
}

type mockNetworkManager struct {
	mock.Mock
}

func (m *mockNetworkManager) ListConnections() ([]vpnmanager.Connection, error) {
	args := m.Called()
	return args.Get(0).([]vpnmanager.Connection), args.Error(1)
}

func (m *mockNetworkManager) SetConnection(conn vpnmanager.Connection) error {
	args := m.Called(conn)
	return args.Error(0)
}
