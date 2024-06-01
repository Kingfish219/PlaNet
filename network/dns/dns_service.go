package dns

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/Kingfish219/PlaNet/network/ni"
)

type DnsService struct {
	PreviousIPConfiguration ni.IPConfiguration
}

type DnsOperation int

const (
	SetDns DnsOperation = iota
	ResetDns
)

func (dnsService DnsService) ChangeDns(operation DnsOperation, dns Dns) (bool, error) {
	activeInterfaceNames, err := ni.GetActiveNetworkInterface()
	if activeInterfaceNames == nil || err != nil {
		return false, err
	}

	if len(activeInterfaceNames) == 0 {
		return false, errors.New("no active network interface found")
	}

	currentIpConfig := ni.IPConfiguration{}
	var result bool

	if operation == SetDns {
		currentIpConfig.IPAddress = ""
		currentIpConfig.SubnetMask = ""
		currentIpConfig.DefaultGateway = ""

		for _, activeInterfaceName := range activeInterfaceNames {
			// result, ipConfig, _ := getStaticIPConfiguration(activeInterfaceName)
			// if result && ipConfig != nil {
			// 	fmt.Println(ipConfig)
			// 	currentIpConfig = *ipConfig
			// }

			result, err = dnsService.setDns(activeInterfaceName, dns.PrimaryDns, dns.SecendaryDns)
		}

		return result, err
	} else {
		for _, activeInterfaceName := range activeInterfaceNames {
			result, err = dnsService.resetDns(activeInterfaceName)
		}

		// if result && currentIpConfig.IPAddress != "" {
		// 	fmt.Println("Reset")
		// 	fmt.Println(currentIpConfig)
		// 	setStaticIPConfiguration(activeInterfaceName, &currentIpConfig)
		// }
		return result, err
	}
}

func (dnsService DnsService) resetDns(interfaceName string) (bool, error) {
	cmd := exec.Command("netsh", "interface", "ipv4", "set", "dnsservers", "name="+interfaceName, "source=dhcp")
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dnsService DnsService) setDns(interfaceName string, dns1 string, dns2 string) (bool, error) {
	commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`,
		interfaceName, dns1, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to execute command: %s, output: %s", err, output)
	}

	return true, nil
}
