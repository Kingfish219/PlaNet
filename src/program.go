package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {

	var activeInterface = getActiveNetworkInterface()
	if activeInterface == "" {
		fmt.Print("Failed getting active network interface")
		return
	}

	dns1 := "8.8.8.8"
	dns2 := "8.8.4.4"

	commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`,
		activeInterface, dns1, dns2)
	//commandText := fmt.Sprintf(`ping "%s"`, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	// Execute the PowerShell command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s, Output: %s\n", err, output)
		return
	}

	fmt.Println(output)
	fmt.Println("DNS settings updated successfully.")
}

func getActiveNetworkInterface() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return ""
	}

	for _, intf := range interfaces {
		if intf.Flags&net.FlagUp != 0 && !strings.Contains(intf.Flags.String(), "loopback") {
			return intf.Name
		}
	}

	return ""
}
