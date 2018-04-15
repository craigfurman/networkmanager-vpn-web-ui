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

	httpHandler := vpnmanager.NewAPI(networkManager)
	server := httptest.NewServer(httpHandler)
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

type mockNetworkManager struct {
	mock.Mock
}

func (m *mockNetworkManager) ListConnections() ([]vpnmanager.Connection, error) {
	args := m.Called()
	return args.Get(0).([]vpnmanager.Connection), args.Error(1)
}
