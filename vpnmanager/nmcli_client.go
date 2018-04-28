package vpnmanager

import (
	"fmt"
	"os/exec"
	"strings"
)

type NmcliClient struct{}

func (*NmcliClient) ListConnections() ([]Connection, error) {
	output, err := exec.Command("nmcli", "connection", "show", "--order", "name").CombinedOutput()
	if err != nil {
		// TODO test
		return nil, err
	}
	return ParseVPNConnectionList(string(output)), nil
}

func (*NmcliClient) SetConnection(conn Connection) error {
	verb := "down"
	if conn.Active {
		verb = "up"
	}

	output, err := exec.Command("nmcli", "connection", verb, conn.Name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("couldn't set connection: output: '%s', error: '%s'", string(output), err)
	}
	return nil
}

func ParseVPNConnectionList(nmcliConShowOutput string) []Connection {
	var conns []Connection
	lines := strings.Split(nmcliConShowOutput, "\n")
	for idx, line := range lines {
		if idx == 0 || line == "" {
			continue
		}

		cols := strings.Fields(line)
		if cols[2] == "vpn" {
			conns = append(conns, Connection{
				Name: cols[0], Active: cols[3] != "--",
			})
		}
	}
	return conns
}
