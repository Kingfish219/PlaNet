package network

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/Kingfish219/PlaNet/internal/domain"
)

func getStaticIPConfiguration(interfaceName string) (bool, *domain.IPConfiguration, error) {
	cmd := exec.Command("netsh", "interface", "ipv4", "show", "config", "name="+interfaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	ipConfig := &domain.IPConfiguration{}
	read := false

	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if strings.Contains(line, strings.ToLower("Configuration for interface")) {
			read = true
			continue
		}

		if read && strings.HasPrefix(line, strings.ToLower("DHCP enabled")) {
			fields := strings.Fields(line)
			var enabled = fields[2]
			if strings.EqualFold(enabled, "yes") {
				return true, nil, nil
			}
		}

		if read && strings.HasPrefix(line, strings.ToLower("IP Address")) {
			fields := strings.Fields(line)
			ipConfig.IPAddress = fields[2]
		}

		if read && strings.HasPrefix(line, strings.ToLower("Subnet")) {
			fields := strings.Fields(line)
			ipConfig.SubnetMask = strings.TrimRight(fields[4], ")")
		}

		if read && strings.HasPrefix(line, strings.ToLower("Default Gateway")) {
			fields := strings.Fields(line)
			ipConfig.DefaultGateway = fields[2]
		}
	}

	if ipConfig.IPAddress == "" || ipConfig.SubnetMask == "" {
		return false, ipConfig, fmt.Errorf("no static IP configuration found for interface: %s", interfaceName)
	}

	return true, ipConfig, nil
}

func setStaticIPConfiguration(interfaceName string, ipConfig *domain.IPConfiguration) error {
	cmd := exec.Command("netsh", "interface", "ipv4", "set", "address", "name="+interfaceName,
		"source=static", "addr="+ipConfig.IPAddress, "mask="+ipConfig.SubnetMask,
		"gateway="+ipConfig.DefaultGateway)
	return cmd.Run()
}

func getActiveNetworkInterface() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return nil, err
	}

	activeInterfaceNames := make([]string, len(interfaces))

	for _, intf := range interfaces {
		if intf.Flags&net.FlagUp != 0 && !strings.Contains(intf.Flags.String(), "loopback") {
			activeInterfaceNames = append(activeInterfaceNames, intf.Name)
		}
	}

	return activeInterfaceNames, nil
}
