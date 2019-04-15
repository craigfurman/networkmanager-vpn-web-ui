package vpnmanager_test

import (
	"testing"

	"github.com/craigfurman/networkmanager-vpn-web-ui/vpnmanager"
	"github.com/stretchr/testify/assert"
)

const exampleOutput = `NAME                  UUID                                  TYPE  DEVICE
a_wifi_network        bd735415-a69c-408e-8a4f-2fcce1c44d61  wifi  iface
a_vpn [foo|bar] [udp] d689c339-c9de-44bf-b8d4-1069d28b47fe  vpn   --   
another_wifi_network  52500cc8-44f9-4587-bf0c-450d9859492d  wifi  --   
another_vpn           b89368fc-2e50-4767-8429-165bfe5dd3cc  vpn   iface
`

func TestParseVPNConnectionListReturnsVPNConnectionsSortedAlphabetically(t *testing.T) {
	conns := vpnmanager.ParseVPNConnectionList(exampleOutput)
	expectedConns := []vpnmanager.Connection{
		{Name: "a_vpn [foo|bar] [udp]", Active: false},
		{Name: "another_vpn", Active: true},
	}
	assert.Equal(t, expectedConns, conns)
}
