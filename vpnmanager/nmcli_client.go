package vpnmanager

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var vpnConnRe = regexp.MustCompile(`^(?P<name>.+)\s+\w{8}\-\w{4}\-\w{4}\-\w{4}\-\w{12}\s+(?P<type>\w+)\s+(?P<active>.*)$`)

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

		captures := vpnConnRe.FindStringSubmatch(line)
		if strings.TrimSpace(captures[2]) == "vpn" {
			conns = append(conns, Connection{
				Name: strings.TrimSpace(captures[1]), Active: strings.TrimSpace(captures[3]) != "--",
			})
		}
	}
	return conns
}
